package awscron

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		expr     string
		expected CronBreakdown
	}{
		{
			"cron(0 18 ? * MON-FRI *)",
			CronBreakdown{
				minutes:   []CronValue{{fieldVal: 0}},
				hours:     []CronValue{{fieldVal: 18}},
				dayMonths: []CronValue{{fieldVal: Unset}},
				months:    []CronValue{{fieldVal: Wildcard}},
				dayWeeks:  []CronValue{{fieldVal: 2, postSepFieldVal: 6, sep: '-'}},
				years:     []CronValue{{fieldVal: Wildcard}},
			},
		},
		{
			"cron(0 9 2#1 * ? 2007)",
			CronBreakdown{
				minutes:   []CronValue{{fieldVal: 0}},
				hours:     []CronValue{{fieldVal: 9}},
				dayMonths: []CronValue{{fieldVal: 2, postSepFieldVal: 1, sep: '#'}},
				months:    []CronValue{{fieldVal: Wildcard}},
				dayWeeks:  []CronValue{{fieldVal: Unset}},
				years:     []CronValue{{fieldVal: 2007}},
			},
		},
	}

	for _, tc := range testCases {
		result, err := format.Parse(tc.expr)

		breakdown := tc.expected

		assert.Equal(0, len(err))
		assert.Equal(
			breakdown,
			result,
			"it returns a breakdown of the expression",
		)
	}
}
