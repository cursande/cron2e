package cron2e

import (
	"errors"
)

type PredefinedCronParser struct {
	expr string
}

// TODO: Handling @reboot, as well as @yearly/annually?
func (parser *PredefinedCronParser) parse() (cb *CronBreakdown, parseErr error) {
	switch parser.expr {
	case "@monthly":
		return &CronBreakdown{
			minutes:   []CronValue{{fieldVal: 0}},
			hours:     []CronValue{{fieldVal: 0}},
			dayMonths: []CronValue{{fieldVal: 1}},
			months:    []CronValue{{fieldVal: 1}},
			dayWeeks:  []CronValue{{fieldVal: Wildcard}},
		}, nil
	case "@weekly":
		return &CronBreakdown{
			minutes:   []CronValue{{fieldVal: 0}},
			hours:     []CronValue{{fieldVal: 0}},
			dayMonths: []CronValue{{fieldVal: Wildcard}},
			months:    []CronValue{{fieldVal: Wildcard}},
			dayWeeks:  []CronValue{{fieldVal: 0}},
		}, nil
	case "@daily":
		return &CronBreakdown{
			minutes:   []CronValue{{fieldVal: 0}},
			hours:     []CronValue{{fieldVal: 0}},
			dayMonths: []CronValue{{fieldVal: Wildcard}},
			months:    []CronValue{{fieldVal: Wildcard}},
			dayWeeks:  []CronValue{{fieldVal: Wildcard}},
		}, nil
	case "@hourly":
		return &CronBreakdown{
			minutes:   []CronValue{{fieldVal: 0}},
			hours:     []CronValue{{fieldVal: Wildcard}},
			dayMonths: []CronValue{{fieldVal: Wildcard}},
			months:    []CronValue{{fieldVal: Wildcard}},
			dayWeeks:  []CronValue{{fieldVal: Wildcard}},
		}, nil
	}

	return nil, errors.New("Unknown predefined scheduling definition")
}
