package cron2e

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	awscron "github.com/cursande/cron2e/src/aws_cron"
	awsrate "github.com/cursande/cron2e/src/aws_rate"
	predefined "github.com/cursande/cron2e/src/predefined"
	standard "github.com/cursande/cron2e/src/standard"
)

type Format interface {
	Translate(expr string) (translation string, errs []error)
}

const (
	standardCron    string = `((\*|((?i)jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)|(\d+|((?i)mon|tue|wed|thur|fri|sat|sun))((\/|\-){0,1}(\d+|((?i)mon|tue|wed|thur|fri|sat|sun)|((?i)jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec))))?){5}`
	predefinedCron  string = `@(annually|yearly|monthly|weekly|daily|hourly|reboot)`
	awsStandardCron string = `cron\(((\*|\?|(\d+|((?i)mon|tue|wed|thur|fri|sat|sun))((\/|\-|\#){0,1}(\d+|((?i)mon|tue|wed|thur|fri|sat|sun)))*)\s*){6}\)`
	awsRateCron     string = `rate\(\d+ (minute(s)?|hour(s)?|day(s)?)\)`
)

func matchCronFormat(pattern string, expr string) bool {
	return regexp.MustCompile(pattern).MatchString(expr)
}

func FormatForExpression(expr string) (Format, error) {
	if strings.Contains(expr, "cron(") || strings.Contains(expr, "rate(") {
		if matchCronFormat(awsStandardCron, expr) {
			return &awscron.AWSCronFormat{}, nil
		}

		if matchCronFormat(awsRateCron, expr) {
			return &awsrate.AWSRateFormat{}, nil
		}
	} else {
		if matchCronFormat(standardCron, expr) {
			return &standard.StandardFormat{}, nil
		}

		if matchCronFormat(predefinedCron, expr) {
			return &predefined.PredefinedFormat{}, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Unrecognised or invalid cron format: '%s'", expr))
}
