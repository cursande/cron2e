package cron2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPredefinedCronParser(t *testing.T) {
	assert := assert.New(t)

	expr := "@monthly"
	parser := &PredefinedCronParser{expr: expr}
	result, err := parser.parse()

	breakdown := &CronBreakdown{
		minutes:   []CronValue{{fieldVal: 0}},
		hours:     []CronValue{{fieldVal: 0}},
		dayMonths: []CronValue{{fieldVal: 1}},
		months:    []CronValue{{fieldVal: 1}},
		dayWeeks:  []CronValue{{fieldVal: Wildcard}},
	}

	assert.Equal(err, nil)

	assert.Equal(
		breakdown,
		result,
		"it returns a breakdown of the expression",
	)
}
