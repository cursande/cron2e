package cron2e

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// - 'Standard'   e.g. "5 4 * * *"
// - 'Predefined' e.g. "@yearly"
// - 'Interval'   e.g. "@every 12h"

// TODO:
// https://golangbyexample.com/runtime-polymorphism-go/
// https://www.golang-book.com/books/intro/9

type Parser interface {
	parse()
}

type StandardCronParser struct {
	expr string
	err  error
}

type PredefinedCronParser struct {
	expr string
	err  error
}

type IntervalCronParser struct {
	expr string
	err  error
}

type AWSCronParser struct {
	expr string
	err  error
}

type AWSRateParser struct {
	expr string
	err  error
}

// TODO: Find a better way to manage flow here
// TODO: in each case where we match digits, should really validate that those digits are actually
// within range for their relevant values. Probably most sensible to validate this in the parser itself,
// rather than with the regex itself
func FindParserForExpression(expr string) (parser interface{}, err error) {
	if strings.Contains(expr, "cron(") || strings.Contains(expr, "rate(") {
		AWSCronR := regexp.MustCompile(`cron\(((\*|\?|\d+((\/|\-){0,1}(\d+))*)\s*){6}\)`)
		if AWSCronR.MatchString(expr) {
			return &AWSCronParser{expr: expr}, nil
		}

		AWSRateR := regexp.MustCompile(`rate\(\d+ (minute(s)?|hour(s)?|day(s)?)\)`)
		if AWSRateR.MatchString(expr) {
			return &AWSRateParser{expr: expr}, nil
		}
	} else {
		// FIXME: need to break below into individual patterns for days, months etc. Will be gross though
		standardCronR := regexp.MustCompile(`(((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|((?i)mon|tue|wed|thur|fri|sat|sun)|((?i)jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)|\*) ?){5,7}`)
		if standardCronR.MatchString(expr) {
			return &StandardCronParser{expr: expr}, nil
		}

		PredefinedCronR := regexp.MustCompile(`@(annually|yearly|monthly|weekly|daily|hourly|reboot)`)
		if PredefinedCronR.MatchString(expr) {
			return &PredefinedCronParser{expr: expr}, nil
		}

		IntervalCronR := regexp.MustCompile(`@every (\d+(ns|us|µs|ms|s|m|h))+`)
		if IntervalCronR.MatchString(expr) {
			return &IntervalCronParser{expr: expr}, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Unrecognised or invalid cron format: %s", expr))
}
