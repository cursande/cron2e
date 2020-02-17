package cron2e

import (
	"errors"
)

type PredefinedCronParser struct {
	expr string
}

func (parser *PredefinedCronParser) parse() (cb *CronBreakdown, parseErr error) {
	return nil, errors.New("not yet implemented")
}
