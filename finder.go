package cron2e

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Parser interface {
	parse(string) (*CronBreakdown, error)
}

type PredefinedCronParser struct {
	expr string
}

type IntervalCronParser struct {
	expr string
}

type AWSCronParser struct {
	expr string
}

type AWSRateParser struct {
	expr string
}

const standardCron string = `(((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|((?i)mon|tue|wed|thur|fri|sat|sun)|((?i)jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)|\*) ?){5,7}`

const predefinedCron string = `@(annually|yearly|monthly|weekly|daily|hourly|reboot)`

const intervalCron string = `@every (\d+(ns|us|µs|ms|s|m|h))+`

const awsStandardCron string = `cron\(((\*|\?|(\d+|((?i)mon|tue|wed|thur|fri|sat|sun))((\/|\-|\#){0,1}(\d+|((?i)mon|tue|wed|thur|fri|sat|sun)))*)\s*){6}\)`

const awsRateCron string = `rate\(\d+ (minute(s)?|hour(s)?|day(s)?)\)`

func matchCronFormat(pattern string, expr string) bool {
	r := regexp.MustCompile(pattern)
	return r.MatchString(expr)
}

func ParserForExpression(expr string) (parser interface{}, err error) {
	if strings.Contains(expr, "cron(") || strings.Contains(expr, "rate(") {
		if matchCronFormat(awsStandardCron, expr) {
			return &AWSCronParser{expr: expr}, nil
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

	return nil, errors.New(fmt.Sprintf("Unrecognised or invalid cron format: %s", expr))
}
