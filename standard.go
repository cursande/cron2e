package cron2e

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type StandardCronParser struct {
	expr string
}

var dayWeekAliases = map[string]int{
	"SUN": 0,
	"MON": 1,
	"TUE": 2,
	"WED": 3,
	"THU": 4,
	"FRI": 5,
	"SAT": 6,
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

func stringValToInt(val string) (newVal int, err error) {
	if val == "*" {
		return Wildcard, nil
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
			return stringValToInt(val)
		}
	} else {
		return stringValToInt(val)
	}

	return
}

func determineAlias(fieldType uint8) (alias map[string]int) {
	switch fieldType {
	case Month:
		return monthAliases
	case DayWeek:
		return dayWeekAliases
	default:
		return nil
	}
}

func buildValueFromPair(pair []string, alias map[string]int, sep rune) (cv CronValue, err error) {
	cv = BuildCronValue()
	cv.sep = sep

	coerced, err := coerceVal(pair[0], alias)
	if err != nil {
		return cv, err
	}
	cv.fieldVal = coerced

	coerced, err = coerceVal(pair[1], alias)
	if err != nil {
		return cv, err
	}
	cv.postSepFieldVal = coerced

	return cv, nil
}

// Coerces a string token into a representation of coerced field values
func tokenToField(token string, fieldType uint8) (cvs []CronValue, err error) {
	alias := determineAlias(fieldType)
	fields := strings.Split(token, ",")

	for i := 0; i < len(fields); i++ {
		field := fields[i]

		if strings.ContainsRune(field, '-') {
			pair := strings.Split(field, "-")

			cv, err := buildValueFromPair(pair, alias, '-')

			if err != nil {
				return nil, err
			}

			cvs = append(cvs, cv)
		} else if strings.ContainsRune(field, '/') {
			pair := strings.Split(field, "/")

			cv, err := buildValueFromPair(pair, alias, '/')

			if err != nil {
				return nil, err
			}

			cvs = append(cvs, cv)
		} else {
			cv := BuildCronValue()

			coerced, err := coerceVal(field, alias)

			if err != nil {
				return nil, err
			}

			cv.fieldVal = coerced
			cvs = append(cvs, cv)
		}
	}

	return cvs, nil
}

func (parser *StandardCronParser) parse() (cb *CronBreakdown, parseErr error) {
	cb = BuildBreakdown()
	tokens := strings.Split(parser.expr, ` `)

	switch len(tokens) {
	case 7:
		return nil, errors.New("not implemented!")
	case 6:
		return nil, errors.New("not implemented!")
	default:
		cb.minutes, parseErr = tokenToField(tokens[0], Minute)
		cb.hours, parseErr = tokenToField(tokens[1], Hour)
		cb.dayMonths, parseErr = tokenToField(tokens[2], DayMonth)
		cb.months, parseErr = tokenToField(tokens[3], Month)
		cb.dayWeeks, parseErr = tokenToField(tokens[4], DayWeek)

		if parseErr != nil {
			return nil, parseErr
		}
	}

	return cb, nil
}
