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
			minutes:   []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
			hours:     []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
			dayMonths: []CronValue{{fieldVal: 1, postSepFieldVal: Unset}},
			months:    []CronValue{{fieldVal: Wildcard, postSepFieldVal: 1, sep: '/'}},
			dayWeeks:  []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
		}, nil
	case "@weekly":
		return &CronBreakdown{
			minutes:   []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
			hours:     []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
			dayMonths: []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
			months:    []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
			dayWeeks:  []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
		}, nil
	case "@daily":
		return &CronBreakdown{
			minutes:   []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
			hours:     []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
			dayMonths: []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
			months:    []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
			dayWeeks:  []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
		}, nil
	case "@hourly":
		return &CronBreakdown{
			minutes:   []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
			hours:     []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
			dayMonths: []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
			months:    []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
			dayWeeks:  []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
		}, nil
	}

	return nil, errors.New("Unknown predefined scheduling definition")
}
