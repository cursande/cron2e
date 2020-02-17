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

func TestValidateFieldVals(t *testing.T) {
	assert := assert.New(t)

	min := CronField{
		fieldVals:        []int{5},
		postSepFieldVals: []int{68},
		sep:              '/',
	}

	valid, err := validateFieldVals(min, validateMinute)

	assert.Equal(false, valid)
	assert.Equal(errors.New("The minute '68' is invalid"), err)
}

func TestValidate(t *testing.T) {
	assert := assert.New(t)

	cb := &CronBreakdown{
		minute: CronField{
			fieldVals:        []int{5},
			postSepFieldVals: []int{65},
		},
		hour: CronField{
			fieldVals:        []int{-1},
			postSepFieldVals: []int{6},
		},
		dayMonth: CronField{
			fieldVals:        []int{-1},
			postSepFieldVals: []int{32},
		},
		month: CronField{
			fieldVals:        []int{6},
			postSepFieldVals: []int{13},
		},
		dayWeek: CronField{
			fieldVals: []int{-1},
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
