package cron2e

import (
	"errors"
)

type AWSRateParser struct {
	expr string
}

func (parser *AWSRateParser) parse() (cb *CronBreakdown, parseErr error) {
	return nil, errors.New("not yet implemented")
}
