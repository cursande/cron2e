package cron2e

import (
	"fmt"
	"errors"
	"regexp"
	"strconv"
)

type AWSRateParser struct {
	expr string
}

func (parser *AWSRateParser) parse() (cb *CronBreakdown, parseErr error) {
	match := regexp.MustCompile(`rate\((\d+) (\w+)\)`).FindAllStringSubmatch(parser.expr, 2)

	num, interval := match[0][1], match[0][2]

	val, err := strconv.Atoi(num)

	if err != nil {
		return nil, err
	}

	switch interval {
	case "minute", "minutes":
		return &CronBreakdown{
			minutes:   []CronValue{{fieldVal: Wildcard, postSepFieldVal: val, sep: '/'}},
			hours:     []CronValue{{fieldVal: Wildcard}},
			dayMonths: []CronValue{{fieldVal: Wildcard}},
			months:    []CronValue{{fieldVal: Wildcard}},
			dayWeeks:  []CronValue{{fieldVal: Wildcard}},
		}, nil
	case "hour", "hours":
		return &CronBreakdown{
			minutes:   []CronValue{{fieldVal: 0}},
			hours:     []CronValue{{fieldVal: Wildcard, postSepFieldVal: val, sep: '/'}},
			dayMonths: []CronValue{{fieldVal: Wildcard}},
			months:    []CronValue{{fieldVal: Wildcard}},
			dayWeeks:  []CronValue{{fieldVal: Wildcard}},
		}, nil
	case "day", "days":
		return &CronBreakdown{
			minutes:   []CronValue{{fieldVal: 0}},
			hours:     []CronValue{{fieldVal: 0}},
			dayMonths: []CronValue{{fieldVal: Wildcard, postSepFieldVal: val, sep: '/'}},
			months:    []CronValue{{fieldVal: Wildcard}},
			dayWeeks:  []CronValue{{fieldVal: Wildcard}},
		}, nil

	default:
		return nil, errors.New(fmt.Sprintf("Cannot parse AWS rate: '%s'", interval))
	}
}
