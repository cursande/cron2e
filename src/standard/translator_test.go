package standard

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

	testCases := []struct {
		expr        string
		expectedRes string
	}{
		{
			"15 4 * 8-9 *",
			"Runs every day from months August through September at 04:15",
		},
		{
			"15 14 1 * *",
			"Runs on the 1st day of the month at 14:15",
		},
		{
			"15 14 1,4 * *",
			"Runs on the 1st and 4th day of the month at 14:15",
		},
		{
			"0 */4 * * *",
			"Runs every day every 4th hour at minute 0",
		},
	}

	for _, tc := range testCases {
		result, errs := format.Translate(tc.expr)

		assert.Equal(0, len(errs))
		assert.Equal(tc.expectedRes, result)
	}
}
