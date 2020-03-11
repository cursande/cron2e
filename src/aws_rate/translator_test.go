package awsrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	assert := assert.New(t)

	cb := &CronBreakdown{
		timeValue: Hour,
		interval:  8,
	}

	result := format.Translate(cb)

	assert.Equal("Runs every 8 hours", result)
}
