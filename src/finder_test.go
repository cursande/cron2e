package cron2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindParserForStandardExpression(t *testing.T) {
	assert := assert.New(t)

	standardTestCases := []string{
		"5 4 * * *",
		"5 0 * 8 *",
		"15 14 1 * *",
		"0 22 * * 1-5",
		"23 0-20/2 * * *",
		"5 4 * * sun",
		"0 0,12 1 */2 *",
		"5 0 * feb-jun *",
		"0 */4 * * *",
	}

	for _, expr := range standardTestCases {
		parser, err := ParserForExpression(expr)

		assert.Equal( nil, err)

		assert.Equal(
			&StandardCronParser{expr: expr},
			parser,
			"it returns the correct parser for a standard cron expression",
		)
	}
}

func TestFindParserForAWSCronExpression(t *testing.T) {
	assert := assert.New(t)

	awsCronTestCases := []string{
		"cron(15 10 * * ? *)",
		"cron(0 8 1 * ? *)",
		"cron(0 18 ? * MON-FRI *)",
		"cron(0 9 ? * 2#1 *)",
	}

	for _, expr := range awsCronTestCases {
		parser, err := ParserForExpression(expr)

		assert.Equal(nil, err)

		assert.Equal(
			&AWSStandardParser{expr: expr},
			parser,
			"it returns the correct parser for an AWS cron expression",
		)
	}
}

func TestFindParserForAWSRateExpression(t *testing.T) {
	assert := assert.New(t)

	awsRateTestCases := []string{
		"rate(5 minutes)",
		"rate(1 hour)",
		"rate(7 days)",
	}

	for _, expr := range awsRateTestCases {
		parser, err := ParserForExpression(expr)

		assert.Equal(nil, err)

		assert.Equal(
			&AWSRateParser{expr: expr},
			parser,
			"it returns the correct parser for an AWS cron expression",
		)
	}
}
