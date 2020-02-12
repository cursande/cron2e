package cron2e

import (
	"errors"
	"fmt"
)

// Will validate a given field's pre and post separator values with a given validator fn.
func validateAllFieldVals(cf CronField, validator func(val int) (bool, error)) (valid bool, validationErr error) {
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

func validateMinuteField(val int) (valid bool, validationErr error) {
	if val < 0 || val > 59 {
		return false, errors.New(fmt.Sprintf("The minute value '%d' is invalid", val))
	}

	return true, nil
}
