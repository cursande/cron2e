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

func (format *PredefinedFormat) Translate(expr string) (translation string, errs []error) {
	bd, found := exprToTrans[expr]

	if !found {
		errs = append(errs, errors.New("Unknown cron expression"))
	}

	return bd, errs
}
