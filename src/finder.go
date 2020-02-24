package cron2e

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type parser interface {
	parse() (*CronBreakdown, error)
}

const standardCron string = `((\*|((?i)jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)|(\d+|((?i)mon|tue|wed|thur|fri|sat|sun))((\/|\-){0,1}(\d+|((?i)mon|tue|wed|thur|fri|sat|sun)|((?i)jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)))?)?){5,7}`

const predefinedCron string = `@(annually|yearly|monthly|weekly|daily|hourly|reboot)`

const intervalCron string = `@every (\d+(ns|us|µs|ms|s|m|h))+`

const awsStandardCron string = `cron\(((\*|\?|(\d+|((?i)mon|tue|wed|thur|fri|sat|sun))((\/|\-|\#){0,1}(\d+|((?i)mon|tue|wed|thur|fri|sat|sun)))*)\s*){6}\)`

const awsRateCron string = `rate\(\d+ (minute(s)?|hour(s)?|day(s)?)\)`

func matchCronFormat(pattern string, expr string) bool {
	r := regexp.MustCompile(pattern)
	return r.MatchString(expr)
}

func ParserForExpression(expr string) (parser interface{parser}, err error) {
	if strings.Contains(expr, "cron(") || strings.Contains(expr, "rate(") {
		if matchCronFormat(awsStandardCron, expr) {
			return &AWSStandardParser{expr: expr}, nil
		}

		if matchCronFormat(awsRateCron, expr) {
			return &AWSRateParser{expr: expr}, nil
		}
	} else {
		if matchCronFormat(standardCron, expr) {
			return &StandardCronParser{expr: expr}, nil
		}

		if matchCronFormat(predefinedCron, expr) {
			return &PredefinedCronParser{expr: expr}, nil
		}

		if matchCronFormat(intervalCron, expr) {
			return &IntervalCronParser{expr: expr}, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Unrecognised or invalid cron format: '%s'", expr))
}
