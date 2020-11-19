package awsrate

import (
	"fmt"
	"strconv"
	"strings"
)

var valToSingularStr = map[uint8]string{
	Minute: "minute",
	Hour:   "hour",
	Day:    "day",
}

var valToPluralStr = map[uint8]string{
	Minute: "minutes",
	Hour:   "hours",
	Day:    "days",
}

func breakdownToStr(timeValue uint8, interval int) string {
	if interval == 1 {
		valToStr := valToSingularStr[timeValue]
		intervalToStr := ""

		return strings.Join([]string{intervalToStr, valToStr}, " ")
	} else {
		valToStr := valToPluralStr[timeValue]
		intervalToStr := strconv.Itoa(interval)

		return strings.Join([]string{intervalToStr, valToStr}, " ")
	}
}

func (format AWSRateFormat) Translate(expr string) (translation string, errs []error) {
	breakdown, errs := format.Parse(expr)

	if len(errs) > 0 {
		return translation, errs
	}

	return fmt.Sprintf("Runs every %s", breakdownToStr(breakdown.timeValue, breakdown.interval)), errs
}
