package cron2e

import (
	"fmt"
	"errors"
	"strings"
	"strconv"
	"regexp"
)

// FIXME: should really be 'namespaced' in its own internal package, not enough shared
// logic across parsers

// https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html
type AWSStandardParser struct {
	expr string
}

const requiredFields int = 6

// [SUN-SAT] = 1-7 in AWS, so we need to adjust it for our breakdown where [SUN-SAT] = 0-6
var awsStdDayWeekAliases = map[string]int{
	"SUN": 1,
	"MON": 2,
	"TUE": 3,
	"WED": 4,
	"THU": 5,
	"FRI": 6,
	"SAT": 7,
}

var awsStdMonthAliases = map[string]int{
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
func awsStdStringValToInt(val string) (newVal int, err error) {
	if val == "*" {
		return Wildcard, nil
	} else if val == "?" {
		return Unset, nil
	} else {
		return strconv.Atoi(val)
	}
}

func awsStdDetermineAlias(fieldType uint8) (alias map[string]int) {
	switch fieldType {
	case Month:
		return awsStdMonthAliases
	case DayWeek:
		return awsStdDayWeekAliases
	default:
		return nil
	}
}

func awsStdCoerceVal(val string, alias map[string]int, fieldType uint8) (newVal int, err error) {
	if len(alias) != 0 {
		r := regexp.MustCompile(`[a-z|A-Z]`)

		if r.MatchString(val) {
			newVal, found := alias[strings.ToUpper(val)]

			if !found {
				return 0, errors.New(fmt.Sprintf("value not recognised: '%s'", val))
			}

			if fieldType == DayWeek {
				return newVal - 1, nil
			} else {
				return newVal, nil
			}
		} else {
			newVal, err = awsStdStringValToInt(val)
		}
	} else {
		newVal, err = awsStdStringValToInt(val)
	}

	if err != nil {
		return 0, err
	}

	if newVal == Unset {
		return newVal, nil
	}

	if fieldType == DayWeek {
		return newVal - 1, nil
	} else {
		return newVal, nil
	}
}

func awsStdBuildValueFromPair(pair []string, alias map[string]int, sep rune, fieldType uint8) (cv CronValue, err error) {
	cv = BuildCronValue()
	cv.sep = sep

	coerced, err := awsStdCoerceVal(pair[0], alias, fieldType)
	if err != nil {
		return cv, err
	}
	cv.fieldVal = coerced

	coerced, err = awsStdCoerceVal(pair[1], alias, fieldType)
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

	cv = BuildCronValue()
	cv.sep = sep

	coerced, err := awsStdCoerceVal(pair[0], alias, fieldType)
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

func awsStdTokenToField(token string, fieldType uint8) (cvs []CronValue, err error) {
	alias := awsStdDetermineAlias(fieldType)
	fields := strings.Split(token, ",")

	for i := 0; i < len(fields); i++ {
		field := fields[i]

		if strings.ContainsRune(field, '-') {
			pair := strings.Split(field, "-")

			cv, err := awsStdBuildValueFromPair(pair, alias, '-', fieldType)

			if err != nil {
				return nil, err
			}

			cvs = append(cvs, cv)
		} else if strings.ContainsRune(field, '/') {
			pair := strings.Split(field, "/")

			cv, err := awsStdBuildValueFromPair(pair, alias, '/', fieldType)

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
			cv := BuildCronValue()

			coerced, err := awsStdCoerceVal(field, alias, fieldType)

			if err != nil {
				return nil, err
			}

			cv.fieldVal = coerced
			cv.postSepFieldVal = Unset

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

func (parser *AWSStandardParser) parse() (cb *CronBreakdown, parseErr error) {
	cb = BuildBreakdown()

	if MultipleOrZeroQuestionMarks(parser.expr) {
		return nil, errors.New("'?' should only be used to mark day-of-week or day-of-month as unused for AWS Cron")
	}

	tokens := strings.Split(parser.expr[5:len(parser.expr) - 1], " ")

	if len(tokens) < requiredFields {
		return nil, errors.New("Invalid AWS cron expression, not enough values provided")
	}

	cb.minutes, parseErr = awsStdTokenToField(tokens[0], Minute)
	if parseErr != nil {
		return nil, parseErr
	}
	cb.hours, parseErr = awsStdTokenToField(tokens[1], Hour)
	if parseErr != nil {
		return nil, parseErr
	}
	cb.dayMonths, parseErr = awsStdTokenToField(tokens[2], DayMonth)
	if parseErr != nil {
		return nil, parseErr
	}
	cb.months, parseErr = awsStdTokenToField(tokens[3], Month)
	if parseErr != nil {
		return nil, parseErr
	}
	cb.dayWeeks, parseErr = awsStdTokenToField(tokens[4], DayWeek)
	if parseErr != nil {
		return nil, parseErr
	}
	cb.years, parseErr = awsStdTokenToField(tokens[5], Year)
	if parseErr != nil {
		return nil, parseErr
	}

	return cb, nil
}
