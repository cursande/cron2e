package awsrate

import (
	"fmt"
	"errors"
	"regexp"
	"strconv"
)

func (format *AWSRateFormat) Parse(expr string) (cb *CronBreakdown, parseErr error) {
	match := regexp.MustCompile(`rate\((\d+) (\w+)\)`).FindAllStringSubmatch(expr, 2)

	num, interval := match[0][1], match[0][2]

	val, err := strconv.Atoi(num)

	if err != nil {
		return nil, err
	}

	switch interval {
	case "minute", "minutes":
		return &CronBreakdown{timeValue: Minute, interval:  val}, nil
	case "hour", "hours":
		return &CronBreakdown{timeValue: Hour, interval: val}, nil
	case "day", "days":
		return &CronBreakdown{timeValue: Day, interval: val}, nil

	default:
		return nil, errors.New(fmt.Sprintf("Cannot parse AWS rate: '%s'", interval))
	}
}
