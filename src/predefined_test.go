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
		minutes:   []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
		hours:     []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
		dayMonths: []CronValue{{fieldVal: 1, postSepFieldVal: Unset}},
		months:    []CronValue{{fieldVal: Wildcard, postSepFieldVal: 1, sep: '/'}},
		dayWeeks:  []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
	}

	assert.Equal(err, nil)

	assert.Equal(
		breakdown,
		result,
		"it returns a breakdown of the expression",
	)
}
