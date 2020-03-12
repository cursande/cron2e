package awsrate

import (
	"fmt"
	"errors"
	"regexp"
	"strconv"
)

func (format *AWSRateFormat) Parse(expr string) (cb *CronBreakdown, parseErrs []error) {
	match := regexp.MustCompile(`rate\((\d+) (\w+)\)`).FindAllStringSubmatch(expr, 2)

	num, interval := match[0][1], match[0][2]

	val, err := strconv.Atoi(num)

	if err != nil {
		parseErrs = append(parseErrs, err)
		return cb, parseErrs
	}

	switch interval {
	case "minute", "minutes":
		return &CronBreakdown{timeValue: Minute, interval:  val}, nil
	case "hour", "hours":
		return &CronBreakdown{timeValue: Hour, interval: val}, nil
	case "day", "days":
		return &CronBreakdown{timeValue: Day, interval: val}, nil

	default:
		parseErrs = append(parseErrs, errors.New(fmt.Sprintf("Cannot parse AWS rate: '%s'", interval)))
		return nil, parseErrs
	}
}
