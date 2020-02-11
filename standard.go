package cron2e

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"errors"
)

type StandardCronParser struct {
	expr string
}

var dayWeekAliases = map[string]int{
	"SUN": 1,
	"MON": 2,
	"TUE": 3,
	"WED": 4,
	"THU": 5,
	"FRI": 6,
	"SAT": 7,
}

var monthAliases = map[string]int{
	"JAN": 1,
	"FEB": 2,
	"MAR": 3,
	"APR": 4,
	"MAY": 5,
	"JUN": 6,
	"JUL": 7,
	"AUG": 8,
	"SEP": 9,
	"OCT": 10,
	"NOV": 11,
	"DEC": 12,
}

func numorWildcardToInt(val string) (newVal int, err error) {
	if val == "*" {
		return -1, nil
	} else {
		return strconv.Atoi(val)
	}
}

// Coerces an int from string to its appropriate form, using an alias if appropriate e.g. "JUL", "SAT"
func coerceVal(val string, alias map[string]int) (newVal int, err error) {
	if len(alias) != 0 {
		r := regexp.MustCompile(`[a-z|A-Z]`)

		if r.MatchString(val) {
			newVal, found := alias[strings.ToUpper(val)]

			if found {
				return newVal, nil
			}
		} else {
			return numorWildcardToInt(val)
		}
	} else {
		return numorWildcardToInt(val)
	}

	return 0, errors.New(fmt.Sprintf("There was an unknown problem coercing the value '%s'", val))
}

// Coerces a string into a representation of coerced field values
func coerceVals(elts string, alias map[string]int) (newVals []int, err error) {
	if strings.ContainsRune(elts, ',') {
		vals := strings.Split(elts, ",")

		for _, val := range vals {
			coerced, err := coerceVal(val, alias)

			if err != nil {
				return nil, err
			}

			newVals = append(newVals, coerced)
		}

		return newVals, nil
	} else {
		coerced, err := coerceVal(elts, alias)

		if err != nil {
			return nil, err
		} else {
			newVals = append(newVals, coerced)
		}

		return newVals, nil
	}
}

func determineAlias(fieldType int) (alias map[string]int) {
	switch fieldType {
	case 4:
		return monthAliases
	case 5:
		return dayWeekAliases
	default:
		return nil
	}
}

func convertAndSetField(expr string, fieldType int) (fieldVals []int, postSepFieldVals []int, sep rune, err error) {
	alias := determineAlias(fieldType)

	if strings.ContainsRune(expr, '/') {
		elts := strings.Split(expr, "/")

		first, err := coerceVals(elts[0], alias)

		if err != nil {
			fieldVals = first
		}

		second, err := coerceVals(elts[1], alias)

		if err != nil {
			postSepFieldVals = second
		}

		return fieldVals, postSepFieldVals, '/', nil

	} else if strings.ContainsRune(expr, '-') {
		elts := strings.Split(expr, "-")

		first, err := coerceVals(elts[0], alias)

		if err != nil {
			fieldVals = first
		}

		second, err := coerceVals(elts[1], alias)

		if err != nil {
			postSepFieldVals = second
		}

		return fieldVals, postSepFieldVals, '-', nil

	} else {
		fieldVals, err = coerceVals(expr, alias)

		return fieldVals, nil, 0, nil
	}

	return
}

// field types:
// 1 - Minute
// 2 - Hour
// 3 - DayMonth
// 4 - Month
// 5 - DayWeek
// When integrated, 0 and 6 will be seconds and years respectively.
func (parser *StandardCronParser) parse() (cb *CronBreakdown, parseErr error) {
	cb = BuildBreakdown()
	tokens := strings.Split(parser.expr, ` `)

	switch len(tokens) {
	case 7:
		return nil, errors.New("not implemented!")
	case 6:
		return nil, errors.New("not implemented!")
	default:
		// TODO: Clean this up
		minutes := BuildCronField()
		minutes.fieldVals, minutes.postSepFieldVals, minutes.sep, parseErr = convertAndSetField(tokens[0], 1)

		hours := BuildCronField()
		hours.fieldVals, hours.postSepFieldVals, hours.sep, parseErr = convertAndSetField(tokens[1], 2)

		dayMonths := BuildCronField()
		dayMonths.fieldVals, dayMonths.postSepFieldVals, dayMonths.sep, parseErr = convertAndSetField(tokens[2], 3)

		months := BuildCronField()
		months.fieldVals, months.postSepFieldVals, months.sep, parseErr = convertAndSetField(tokens[3], 4)

		dayWeeks := BuildCronField()
		dayWeeks.fieldVals, dayWeeks.postSepFieldVals, dayWeeks.sep, parseErr = convertAndSetField(tokens[4], 5)

		cb.minute = minutes
		cb.hour = hours
		cb.dayMonth = dayMonths
		cb.dayWeek = dayWeeks
		cb.month = months
	}

	if parseErr != nil {
		return nil, parseErr
	}

	return cb, nil
}
