package cron2e

import (
	"errors"
	"fmt"
)

func errorMessage(fieldName string, val int) error {
	return errors.New(fmt.Sprintf("The %s '%d' is invalid", fieldName ,val))
}

func validateMinute(val int) (valid bool, validationErr error) {
	if val <= -1 || val < 60 {
		return true, nil
	}

	return false, errorMessage("minute", val)
}

func validateHour(val int) (valid bool, validationErr error) {
	if val <= -1 || val < 24 {
		return true, nil
	}

	return false, errorMessage("hour", val)
}

func validateDayMonth(val int) (valid bool, validationErr error) {
	if (val > 0 && val < 32) || val == -1 {
		return true, nil
	}

	return false, errorMessage("month-day", val)
}

func validateMonth(val int) (valid bool, validationErr error) {
	if (val > 0 && val < 13) || val == -1 {
		return true, nil
	}

	return false, errorMessage("month", val)
}

// Valid values can vary, depending on whether weekdays are zero-indexed. We assume the standard format
// is 0-6 (Sunday to Saturday)
func validateDayWeek(val int) (valid bool, validationErr error) {
	if val >= -1 || val < 7 {
		return true, nil
	}

	return false, errorMessage("weekday", val)
}

func validateFieldVals(cf CronField, validator func(int) (bool, error)) (valid bool, validationErr error) {
	for i := 0; i < len(cf.fieldVals); i++ {
		val := cf.fieldVals[i]

		valid, err := validator(val)

		if !valid {
			return false, err
		}
	}

	for i := 0; i < len(cf.postSepFieldVals); i++ {
		val := cf.postSepFieldVals[i]

		valid, err := validator(val)

		if !valid {
			return false, err
		}
	}

	return true, nil
}

func Validate(cb *CronBreakdown) (isValid bool) {
	validations := []struct {
		field     CronField
		validator func(int) (bool, error)
	}{
		{cb.minute, validateMinute},
		{cb.hour, validateHour},
		{cb.dayMonth, validateDayMonth},
		{cb.month, validateMonth},
		{cb.dayWeek, validateDayWeek},
	}

	for _, v := range validations {
		valid, err := validateFieldVals(v.field, v.validator)

		if !valid {
			cb.validationErrs = append(cb.validationErrs, err)
		}
	}

	if len(cb.validationErrs) > 0 {
		return false
	}

	return true
}
