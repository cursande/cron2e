package standard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardCronParser(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		expr     string
		expected CronBreakdown
	}{
		{
			"5 0 * 8 *",
			CronBreakdown{
				minutes:   []CronValue{{fieldVal: 5, postSepFieldVal: Unset}},
				hours:     []CronValue{{fieldVal: 0, postSepFieldVal: Unset}},
				dayMonths: []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
				months:    []CronValue{{fieldVal: 8, postSepFieldVal: Unset}},
				dayWeeks:  []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
			},
		},
		{
			"15 4 * 8-9 *",
			CronBreakdown{
				minutes:   []CronValue{{fieldVal: 15, postSepFieldVal: Unset}},
				hours:     []CronValue{{fieldVal: 4, postSepFieldVal: Unset}},
				dayMonths: []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
				months:    []CronValue{{fieldVal: 8, postSepFieldVal: 9, sep: '-'}},
				dayWeeks:  []CronValue{{fieldVal: Wildcard, postSepFieldVal: Unset}},
			},
		},
	}

	for _, tc := range testCases {
		result, errs := (format.Parse(tc.expr))
		breakdown := tc.expected

		assert.Equal(0, (len(errs)))
		assert.Equal(
			breakdown,
			result,
			"it returns a breakdown of the expression",
		)
	}
}
