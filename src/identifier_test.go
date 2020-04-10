package cron2e

import (
	"testing"

	"github.com/stretchr/testify/assert"

	awscron "github.com/cursande/cron2e/src/aws_cron"
	awsrate "github.com/cursande/cron2e/src/aws_rate"
	standard "github.com/cursande/cron2e/src/standard"
)

func TestFindParserForStandardExpression(t *testing.T) {
	assert := assert.New(t)

	standardTestCases := []string{
		"5 4 * * *",
		"5 0 * 8 *",
		"15 14 1 * *",
		"0 22 * * 1-5",
		"5 4 * * sun",
		"0 0,12 1 */2 *",
		"5 0 * feb-jun *",
		"0 */4 * * *",
	}

	for _, expr := range standardTestCases {
		format, err := FormatForExpression(expr)

		assert.Equal( nil, err)

		assert.Equal(
			&standard.StandardFormat{},
			format,
			"it returns the correct format for a standard cron expression",
		)
	}
}

func TestFindParserForAWSCronExpression(t *testing.T) {
	assert := assert.New(t)

	testCases := []string{
		"cron(15 10 * * ? *)",
		"cron(0 8 1 * ? *)",
		"cron(0 18 ? * MON-FRI *)",
		"cron(0 9 ? * 2#1 *)",
	}

	for _, expr := range testCases {
		format, err := FormatForExpression(expr)

		assert.Equal(nil, err)

		assert.Equal(
			&awscron.AWSCronFormat{},
			format,
			"it returns the correct format for an AWS cron expression",
		)
	}
}

func TestFindParserForAWSRateExpression(t *testing.T) {
	assert := assert.New(t)

	testCases := []string{
		"rate(5 minutes)",
		"rate(1 hour)",
		"rate(7 days)",
	}

	for _, expr := range testCases {
		format, err := FormatForExpression(expr)

		assert.Equal(nil, err)

		assert.Equal(
			&awsrate.AWSRateFormat{},
			format,
			"it returns the correct format for an AWS cron expression",
		)
	}
}
