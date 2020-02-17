package cron2e

import (
	"errors"
)

type IntervalCronParser struct {
	expr string
}

func (parser *IntervalCronParser) parse() (cb *CronBreakdown, parseErr error) {
	return nil, errors.New("not yet implemented")
}
