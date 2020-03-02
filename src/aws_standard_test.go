package cron2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAWSStandardParser(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		expr     string
		expected *CronBreakdown
	}{
		{
			"cron(0 18 ? * MON-FRI *)",
			&CronBreakdown{
				minutes:   []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
				hours:     []CronValue{{fieldVal: 18, postSepFieldVal: Unset}},
				dayMonths: []CronValue{{fieldVal: Unset, postSepFieldVal: Unset}},
				months:    []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
				dayWeeks:  []CronValue{{fieldVal: 1, postSepFieldVal: 5, sep: '-'}},
				years:     []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
			},
		},
		{
			"cron(0 9 2#1 * ? 2007)",
			&CronBreakdown{
				minutes:   []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
				hours:     []CronValue{{fieldVal: 9, postSepFieldVal: Unset}},
				dayMonths: []CronValue{{fieldVal: 2, postSepFieldVal: 1, sep: '#'}},
				months:    []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
				dayWeeks:  []CronValue{{fieldVal: Unset, postSepFieldVal: Unset}},
				years:     []CronValue{{fieldVal: 2007, postSepFieldVal: Unset}},
			},
		},
	}

	for _, tc := range testCases {
		parser := &AWSStandardParser{expr: tc.expr}

		result, err := parser.parse()

		breakdown := tc.expected

		assert.Equal(nil, err)
		assert.Equal(
			breakdown,
			result,
			"it returns a breakdown of the expression",
		)
	}
}
