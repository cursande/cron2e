package awscron

import (
	"errors"
	"strings"
	"strconv"
)

const requiredFields int = 6

var DayWeekAliases = map[string]int{
	"SUN": 1,
	"MON": 2,
	"TUE": 3,
	"WED": 4,
	"THU": 5,
	"FRI": 6,
	"SAT": 7,
}

var MonthAliases = map[string]int{
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

// TODO: Will need to support `L` and `W`
func stringValToInt(val string) (newVal int, err error) {
	if val == "*" {
		return Wildcard, nil
	} else if val == "?" {
		return Unset, nil
	} else {
		return strconv.Atoi(val)
	}
}

func DetermineAlias(fieldType uint8) (alias map[string]int) {
	switch fieldType {
	case Month:
		return MonthAliases
	case DayWeek:
		return DayWeekAliases
	default:
		return nil
	}
}

func CoerceVal(val string, alias map[string]int) (newVal int, err error) {
	if len(alias) != 0 {
		newVal, found := alias[strings.ToUpper(val)]

		if found {
			return newVal, nil
		}
	}

	return stringValToInt(val)
}

func BuildValueFromPair(pair []string, alias map[string]int, sep rune) (cv CronValue, err error) {
	cv = CronValue{}
	cv.sep = sep

	coerced, err := CoerceVal(pair[0], alias)
	if err != nil {
		return cv, err
	}
	cv.fieldVal = coerced

	coerced, err = CoerceVal(pair[1], alias)
	if err != nil {
		return cv, err
	}
	cv.postSepFieldVal = coerced

	return cv, nil
}

// This is a unique concept in AWS crons, the latter value does not really represent a time interval but
func BuildInstanceOfValueFromPair(pair []string, alias map[string]int, sep rune, fieldType uint8) (cv CronValue, err error) {
	if fieldType != DayMonth {
		return cv, errors.New("'#' can only be used for the day-month value in AWS Cron")
	}

	cv = CronValue{}
	cv.sep = sep

	coerced, err := CoerceVal(pair[0], alias)
	if err != nil {
		return cv, err
	}
	cv.fieldVal = coerced

	coerced, err = strconv.Atoi(pair[1])
	if err != nil {
		return cv, err
	}
	cv.postSepFieldVal = coerced

	return cv, nil
}

func TokenToField(token string, fieldType uint8) (cvs []CronValue, err error) {
	alias := DetermineAlias(fieldType)
	fields := strings.Split(token, ",")

	for i := 0; i < len(fields); i++ {
		field := fields[i]

		if strings.ContainsRune(field, '-') {
			pair := strings.Split(field, "-")

			cv, err := BuildValueFromPair(pair, alias, '-')

			if err != nil {
				return nil, err
			}

			cvs = append(cvs, cv)
		} else if strings.ContainsRune(field, '/') {
			pair := strings.Split(field, "/")

			cv, err := BuildValueFromPair(pair, alias, '/')

			if err != nil {
				return nil, err
			}

			cvs = append(cvs, cv)
		} else if strings.ContainsRune(field, '#') {
			pair := strings.Split(field, "#")

			cv, err := BuildInstanceOfValueFromPair(pair, alias, '#', fieldType)

			if err != nil {
				return nil, err
			}

			cvs = append(cvs, cv)
		} else {
			cv := CronValue{}

			coerced, err := CoerceVal(field, alias)

			if err != nil {
				return nil, err
			}

			cv.fieldVal = coerced

			cvs = append(cvs, cv)
		}
	}

	return cvs, nil
}

func MultipleOrZeroQuestionMarks(expr string) bool {
	occ := 0

	for i := 0; i < len(expr); i++ {
		if expr[i] == '?' {
			occ++
		}
	}

	return (occ >= 2 || occ == 0)
}

func (format AWSCronFormat) Parse(expr string) (cb CronBreakdown, parseErrs []error) {
	cb = CronBreakdown{}

	if MultipleOrZeroQuestionMarks(expr) {
		parseErrs = append(parseErrs, errors.New("'?' should only be used to mark day-of-week or day-of-month as unused for AWS Cron"))
		return cb, parseErrs
	}

	tokens := strings.Split(expr[5:len(expr) - 1], " ")

	if len(tokens) < requiredFields {
		parseErrs = append(parseErrs, errors.New("Invalid AWS cron expression, not enough values provided"))
		return cb, parseErrs
	}

	var parseErr error

	cb.minutes, parseErr = TokenToField(tokens[0], Minute)
	if parseErr != nil {
		parseErrs = append(parseErrs, parseErr)
	}
	cb.hours, parseErr = TokenToField(tokens[1], Hour)
	if parseErr != nil {
		parseErrs = append(parseErrs, parseErr)
	}
	cb.dayMonths, parseErr = TokenToField(tokens[2], DayMonth)
	if parseErr != nil {
		parseErrs = append(parseErrs, parseErr)
	}
	cb.months, parseErr = TokenToField(tokens[3], Month)
	if parseErr != nil {
		parseErrs = append(parseErrs, parseErr)
	}
	cb.dayWeeks, parseErr = TokenToField(tokens[4], DayWeek)
	if parseErr != nil {
		parseErrs = append(parseErrs, parseErr)
	}
	cb.years, parseErr = TokenToField(tokens[5], Year)
	if parseErr != nil {
		parseErrs = append(parseErrs, parseErr)
	}

	return cb, parseErrs
}
