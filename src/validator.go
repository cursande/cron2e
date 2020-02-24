package cron2e

import (
	"errors"
	"fmt"
)

func errorMessage(fieldName string, val int) error {
	return errors.New(fmt.Sprintf("The %s '%d' is invalid", fieldName, val))
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

func validateWildcardNotPostSepFieldVal(cv CronValue) (valid bool, validationErr error) {
	if cv.postSepFieldVal != Wildcard {
		return true, nil
	}

	return false, errors.New("A wildcard cannot be used as a post-separator value")
}

func validateRangeWithoutWildcard(cv CronValue) (valid bool, validationErr error) {
	if !(cv.sep == '-' && cv.fieldVal == Wildcard) {
		return true, nil
	}

	return false, errors.New("A wildcard cannot be used in a range expression")
}

func validatePostSepIsNotZero(cv CronValue) (valid bool, validationErr error) {
	if !(cv.postSepFieldVal == 0 && cv.sep != 0) {
		return true, nil
	}

	return false, errors.New("0 cannot be the post-separator value in a range or step expression")
}

func validateCronValue(cv CronValue, validator func(int) (bool, error)) (valid bool, err error) {
	valid, err = validator(cv.fieldVal)

	if !valid {
		return false, err
	}

	if cv.postSepFieldVal == Unset {
		return true, nil
	}

	valid, err = validator(cv.postSepFieldVal)

	if !valid {
		return false, err
	}

	return true, nil
}

func validateField(cvs []CronValue, validator func(int) (bool, error)) (valid bool, validationErr error) {
	for i := 0; i < len(cvs); i++ {
		valid, err := validateCronValue(cvs[i], validator)

		if !valid {
			return false, err
		}
	}

	return true, nil
}

var genericValidations = []func(CronValue) (bool, error) {
	validateWildcardNotPostSepFieldVal,
	validateRangeWithoutWildcard,
	validatePostSepIsNotZero,
}

func Validate(cb *CronBreakdown) (isValid bool) {
	validations := []struct {
		field     []CronValue
		validator func(int) (bool, error)
	}{
		{cb.minutes, validateMinute},
		{cb.hours, validateHour},
		{cb.dayMonths, validateDayMonth},
		{cb.months, validateMonth},
		{cb.dayWeeks, validateDayWeek},
	}

	for _, v := range validations {

		// higher-level validations for all fields
		for _, gv := range genericValidations {
			for _, cv := range v.field {
				valid, err := gv(cv)

				if !valid {
					cb.validationErrs = append(cb.validationErrs, err)
				}
			}

		}

		// field-specific validations on values
		valid, err := validateField(v.field, v.validator)

		if !valid {
			cb.validationErrs = append(cb.validationErrs, err)
		}
	}

	if len(cb.validationErrs) > 0 {
		return false
	}

	return true
}
