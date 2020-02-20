package cron2e

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type IntervalCronParser struct {
	expr string
}

func (parser *IntervalCronParser) parse() (cb *CronBreakdown, parseErr error) {
	match := regexp.MustCompile(`(\d+)(\w{1,2})`).FindAllStringSubmatch(parser.expr, 2)

	num, interval := match[0][1], match[0][2]

	val, err := strconv.Atoi(num)

	if err != nil {
		return nil, err
	}

	switch interval {
	case "m":
		return &CronBreakdown{
			minutes:   []CronValue{{fieldVal: Wildcard, postSepFieldVal: val, sep: '/'}},
			hours:     []CronValue{{fieldVal: Wildcard}},
			dayMonths: []CronValue{{fieldVal: Wildcard}},
			months:    []CronValue{{fieldVal: Wildcard}},
			dayWeeks:  []CronValue{{fieldVal: Wildcard}},
		}, nil
	case "h":
		return &CronBreakdown{
			minutes:   []CronValue{{fieldVal: 0}},
			hours:     []CronValue{{fieldVal: val}},
			dayMonths: []CronValue{{fieldVal: Wildcard}},
			months:    []CronValue{{fieldVal: Wildcard}},
			dayWeeks:  []CronValue{{fieldVal: Wildcard}},
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Cannot process interval: '%s'", interval))
	}
}
