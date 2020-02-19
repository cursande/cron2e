package cron2e

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateMinute(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		val           int
		expectedValid bool
		expectedErr   error
	}{
		{59, true, nil},
		{81, false, errors.New("The minute '81' is invalid")},
		{-1, true, nil},
	}

	for _, tc := range testCases {
		valid, err := validateMinute(tc.val)

		assert.Equal(tc.expectedErr, err)
		assert.Equal(tc.expectedValid, valid)
	}
}

func TestValidateField(t *testing.T) {
	assert := assert.New(t)

	min := []CronValue{
		{
			fieldVal:        5,
			postSepFieldVal: 68,
			sep:             '/',
		},
	}

	valid, err := validateField(min, validateMinute)

	assert.Equal(false, valid)
	assert.Equal(errors.New("The minute '68' is invalid"), err)
}

func TestValidate(t *testing.T) {
	assert := assert.New(t)

	cb := &CronBreakdown{
		minutes: []CronValue{
			{
				fieldVal:        5,
				postSepFieldVal: 65,
				sep:             '-',
			},
		},
		hours: []CronValue{
			{
				fieldVal:        5,
				postSepFieldVal: 10,
				sep:             '-',
			},
		},
		dayMonths: []CronValue{
			{
				fieldVal:        Wildcard,
				postSepFieldVal: 32,
				sep:             '/',
			},
		},
		months: []CronValue{
			{
				fieldVal: 13,
			},
		},
		dayWeeks: []CronValue{
			{
				fieldVal: Wildcard,
			},
		},
	}

	valid := Validate(cb)

	assert.Equal(false, valid)

	assert.Equal(
		[]error{
			errors.New("The minute '65' is invalid"),
			errors.New("The month-day '32' is invalid"),
			errors.New("The month '13' is invalid"),
		},
		cb.validationErrs,
	)
}
