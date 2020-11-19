package awsrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	result, err := format.Parse("rate(5 minutes)")

	breakdown := CronBreakdown{
		timeValue: Minute,
		interval:  5,
	}

	assert.Equal(0, len(err))

	assert.Equal(
		breakdown,
		result,
		"it returns a breakdown of the expression",
	)
}
