package awscron

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateField(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		validator   func(int) (bool, error)
		cvs         []CronValue
		expectedErr error
	}{
		{
			validateMinute,
			[]CronValue{
				{
					fieldVal:        5,
					postSepFieldVal: 68,
					sep:             '/',
				},
			},
			errors.New("The minute '68' is invalid"),
		},


		{
			validateDayWeek,
			[]CronValue{
				{
					fieldVal:        2000,
					postSepFieldVal: Unset,
				},
			},
			errors.New("The weekday '2000' is invalid"),
		},
	}

	for _, tc := range testCases {
		valid, err := validateField(tc.cvs, tc.validator)

		assert.Equal(false, valid)
		assert.Equal(tc.expectedErr, err)
	}
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
				fieldVal:        Wildcard,
				postSepFieldVal: 10,
				sep:             '-',
			},
		},
		dayWeeks: []CronValue{
			{
				fieldVal:        Wildcard,
				postSepFieldVal: Unset,
			},
		},
	}

	errs := format.Validate(cb)

	assert.Equal(
		[]error{
			errors.New("The minute '65' is invalid"),
			errors.New("The month-day '32' is invalid"),
			errors.New("A wildcard cannot be used in a range expression"),
		},
		errs,
	)
}
