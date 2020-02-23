package cron2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAWSRateParser(t *testing.T) {
	assert := assert.New(t)

	expr := "rate(5 minutes)"
	parser := &AWSRateParser{expr: expr}
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
