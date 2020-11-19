package awscron

import (
	"fmt"
	"strconv"
	"strings"
)

var intToDayWeekName = map[int]string{
	1: "Sunday",
	2: "Monday",
	3: "Tuesday",
	4: "Wednesday",
	5: "Thursday",
	6: "Friday",
	7: "Saturday",
}

var intToMonthName = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

var typesToSingularStr = map[uint8]string{
	Minute:   "minute",
	Hour:     "hour",
	DayMonth: "day of the month",
	Month:    "month",
	DayWeek:  "weekday",
}

var typesToPluralStr = map[uint8]string{
	Minute:   "minutes",
	Hour:     "hours",
	DayMonth: "days-of-the-month",
	Month:    "months",
	DayWeek:  "weekdays",
}

func ordinal(val int) string {
	valToS := strconv.Itoa(val)

	ones := val % 10
	tens := val % 100

	suffix := "th"

	if tens < 11 || tens > 13 {
		switch ones {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		}
	}

	return strings.Join([]string{valToS, suffix}, "")
}

func determineTranslationAlias(fieldType uint8) (alias map[int]string) {
	switch fieldType {
	case Month:
		return intToMonthName
	case DayWeek:
		return intToDayWeekName
	default:
		return nil
	}
}

func fieldContainsOnlyWildcard(cvs []CronValue) bool {
	for i := 0; i < len(cvs); i++ {
		if cvs[i].fieldVal == Wildcard && cvs[i].postSepFieldVal == 0 {
			return true
		}
	}

	return false
}

func occursEveryDay(cb CronBreakdown) bool {
	return (fieldContainsOnlyWildcard(cb.dayMonths) || fieldContainsOnlyWildcard(cb.dayWeeks))
}

func isStep(sep rune) bool { return sep == '/' }

func isRange(sep rune) bool { return sep == '-' }

func isInstanceOf(sep rune) bool { return sep == '#' }

func stepToStr(cv CronValue, fieldType uint8) string {
	interval := ordinal(cv.postSepFieldVal)

	if cv.fieldVal == Wildcard {
		if cv.postSepFieldVal == 1 {
			return fmt.Sprintf("every %s", typesToSingularStr[fieldType])
		} else {
			return fmt.Sprintf("every %s %s", ordinal(cv.postSepFieldVal), typesToSingularStr[fieldType])
		}
	} else {
		startRange := valueToStr(cv.fieldVal, fieldType)

		return fmt.Sprintf("every %s%s, starting from %s", interval, typesToSingularStr[fieldType], startRange)
	}

	return fmt.Sprintf("every %s%s", interval, typesToPluralStr[fieldType])
}

func rangeToStr(cv CronValue, fieldType uint8) string {
	from := valueToStr(cv.fieldVal, fieldType)
	through := valueToStr(cv.postSepFieldVal, fieldType)

	return fmt.Sprintf("from %s %s through %s", typesToPluralStr[fieldType], from, through)
}

// Specific to days of the week in AWS
func instanceOfToStr(cv CronValue, fieldType uint8) string {
	weekday := valueToStr(cv.fieldVal, fieldType)
	instance := ordinal(cv.postSepFieldVal)

	return fmt.Sprintf("on the %s %s of the month", instance, weekday)
}

func valueToStr(val int, fieldType uint8) string {
	alias := determineTranslationAlias(fieldType)

	if len(alias) != 0 {
		return alias[val]
	}

	if fieldType == DayMonth {
		return ordinal(val)
	} else {
		return strconv.Itoa(val)
	}
}

func noSeparatorsInField(cvs []CronValue) bool {
	for i := 0; i < len(cvs); i++ {
		if cvs[i].sep != 0 {
			return false
		}
	}

	return true
}

func FieldToStr(cvs []CronValue, fieldType uint8) string {
	parts := []string{}

	for i := 0; i < len(cvs); i++ {
		cv := cvs[i]

		if cv.fieldVal == Unset {
			continue
		}

		if isStep(cv.sep) {
			parts = append(parts, stepToStr(cv, fieldType))

		} else if isRange(cv.sep) {
			parts = append(parts, rangeToStr(cv, fieldType))

		} else if isInstanceOf(cv.sep) {
			parts = append(parts, instanceOfToStr(cv, fieldType))
		} else {
			if cv.fieldVal != Wildcard {
				parts = append(parts, valueToStr(cv.fieldVal, fieldType))
			}
		}
	}

	phrase := strings.Join(parts, " and ")

	if noSeparatorsInField(cvs) && len(parts) > 0 {
		switch fieldType {
		case Month, Year:
			return strings.Join([]string{"in", phrase}, " ")
		case DayMonth:
			return strings.Join([]string{"on the ", phrase, " day of the month"}, "")
		case DayWeek:
			return strings.Join([]string{"on", phrase}, " ")
		case Hour:
			return strings.Join([]string{"at hour", phrase}, " ")
		case Minute:
			return strings.Join([]string{"at minute", phrase}, " ")
		}
	}

	return phrase
}

func canFormatTimeOfDay(minutes []CronValue, hours []CronValue) bool {
	if len(minutes) != 1 || len(hours) != 1 {
		return false
	}

	minute := minutes[0]
	hour := hours[0]

	if minute.sep != 0 || hour.sep != 0 {
		return false
	}

	if minute.fieldVal == Wildcard || hour.fieldVal == Wildcard {
		return false
	}

	return true
}

func combineMinuteAndHour(minutes []CronValue, hours []CronValue) string {
	minute := strconv.Itoa(minutes[0].fieldVal)
	hour := strconv.Itoa(hours[0].fieldVal)

	return (fmt.Sprintf("at %02s:%02s", hour, minute))
}

func removeBlank(segments []string) (res []string) {
	for i := 0; i < len(segments); i++ {
		seg := segments[i]
		if seg != "" {
			res = append(res, seg)
		}
	}

	return res
}

func generateExpression(cb CronBreakdown) string {
	segments := []string{}

	if occursEveryDay(cb) {
		segments = append(segments, "every day")
	} else {
		segments = append(segments, FieldToStr(cb.dayMonths, DayMonth))
		segments = append(segments, FieldToStr(cb.dayWeeks, DayWeek))
	}

	segments = append(segments, FieldToStr(cb.months, Month))

	if canFormatTimeOfDay(cb.minutes, cb.hours) {
		segments = append(segments, combineMinuteAndHour(cb.minutes, cb.hours))
	} else {
		segments = append(segments, FieldToStr(cb.hours, Hour))
		segments = append(segments, FieldToStr(cb.minutes, Minute))
	}

	segments = append(segments, FieldToStr(cb.years, Year))

	translation := strings.Join(removeBlank(segments), " ")

	return fmt.Sprintf("Runs %s", translation)
}

func (format AWSCronFormat) Translate(expr string) (translation string, errs []error) {
	breakdown, errs := format.Parse(expr)

	if len(errs) > 0 {
		return translation, errs
	}

	errs = format.Validate(breakdown)

	if len(errs) > 0 {
		return translation, errs
	}

	return generateExpression(breakdown), nil
}
