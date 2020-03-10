package predefined

import (
	"errors"
)

type PredefinedFormat struct{}

var exprToTrans = map[string]string{
	"@annually": "Runs on January 1 at 00:00",
	"@yearly":   "Runs on January 1 at 00:00",
	"@monthly":  "Runs every month at 00:00",
	"@weekly":   "Runs every Sunday at 00:00",
	"@daily":    "Runs every day at 00:00",
	"@hourly":   "Runs every hour at minute 0",
	"@reboot":   "Runs after cron daemon reboots",
}

// There is no dynamic information that needs to be parsed for a predefined cron.
func (format *PredefinedFormat) Parse(expr string) (bd string, parseErrs []error) {
	bd, found := exprToTrans[expr]

	if !found {
		parseErrs = append(parseErrs, errors.New("Unknown cron expression"))
	}

	return bd, parseErrs
}

func (format *PredefinedFormat) Validate(breakdown string) (validationErrs []error) { return }

func (format *PredefinedFormat) Translate(breakdown string) (translation string) { return breakdown }
