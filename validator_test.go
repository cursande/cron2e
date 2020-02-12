package cron2e

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateMinuteField(t *testing.T) {
	assert := assert.New(t)

	min := CronField{
		fieldVals:        []int{5},
		postSepFieldVals: []int{68},
		sep:              '/',
	}

	valid, err := validateAllFieldVals(min, validateMinuteField)

	assert.Equal(false, valid)
	assert.Equal(errors.New("The minute value '68' is invalid"), err)
}
