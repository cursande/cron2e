package cron2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntervalCronParser(t *testing.T) {
	assert := assert.New(t)

	expr := "@every 5m"
	parser := &IntervalCronParser{expr: expr}
	result, err := parser.parse()

	breakdown := &CronBreakdown{
		minutes:   []CronValue{{fieldVal: Wildcard, postSepFieldVal: 5, sep: '/'}},
		hours:     []CronValue{{fieldVal: Wildcard}},
		dayMonths: []CronValue{{fieldVal: Wildcard}},
		months:    []CronValue{{fieldVal: Wildcard}},
		dayWeeks:  []CronValue{{fieldVal: Wildcard}},
	}

	assert.Equal(err, nil)

	assert.Equal(
		breakdown,
		result,
		"it returns a breakdown of the expression",
	)
}
