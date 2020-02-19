package cron2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardCronParser(t *testing.T) {
	assert := assert.New(t)

	expr := "5 0 * 8 *" // 00:05 in August
	parser := &StandardCronParser{expr: expr}
	result, err := parser.parse()

	breakdown := &CronBreakdown{
		minutes: []CronValue{{fieldVal: 5}},
		hours: []CronValue{{fieldVal: 0}},
		dayMonths: []CronValue{{fieldVal: Wildcard}},
		months: []CronValue{{fieldVal: 8}},
		dayWeeks: []CronValue{{fieldVal: Wildcard}},
	}

	assert.Equal(err, nil)

	assert.Equal(
		breakdown,
		result,
		"it returns a breakdown of the expression",
	)
}
