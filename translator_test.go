package cron2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldToStr(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		cvs         []CronValue
		fieldType   uint8
		expectedStr string
	}{
		{
			[]CronValue{{fieldVal: 8}, {fieldVal: 9}},
			Month,
			"in August and September",
		},
		{
			[]CronValue{{fieldVal: 2}, {fieldVal: Wildcard, postSepFieldVal: 2, sep: '/'}},
			Minute,
			"2 and every 2nd minute",
		},
		{
			[]CronValue{{fieldVal: 8, postSepFieldVal: 9, sep: '-'}},
			Hour,
			"from hours 8 through 9",
		},
		{
			[]CronValue{{fieldVal: Wildcard, postSepFieldVal: 5, sep: '/'}},
			Minute,
			"every 5th minute",
		},
		{
			[]CronValue{{fieldVal: Wildcard, postSepFieldVal: 1, sep: '/'}},
			Hour,
			"every hour",
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			tc.expectedStr,
			FieldToStr(tc.cvs, tc.fieldType),
		)
	}
}

func TestTranslate(t *testing.T) {
	assert := assert.New(t)

	cb := &CronBreakdown{
		minutes:   []CronValue{{fieldVal: 15}},
		hours:     []CronValue{{fieldVal: 4}},
		dayMonths: []CronValue{{fieldVal: Wildcard}},
		months:    []CronValue{{fieldVal: 8, postSepFieldVal: 9, sep: '-'}},
		dayWeeks:  []CronValue{{fieldVal: Wildcard}},
	}

	result, err := Translate(cb)

	assert.Equal(err, nil)
	assert.Equal("Runs from months August through September, at 04:15", result)
}
