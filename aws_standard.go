package cron2e

import (
	"errors"
)

type AWSStandardParser struct {
	expr string
}

func (parser *AWSStandardParser) parse() (cb *CronBreakdown, parseErr error) {
	return nil, errors.New("not yet implemented")
}
