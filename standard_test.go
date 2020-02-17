package cron2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardCronParser(t *testing.T) {
	assert := assert.New(t)

	expr := "5 0 * 8 *" // 00:05 in August
	parser := &StandardCronParser{expr: expr}
	result, err := parser.parse()

	breakdown := &CronBreakdown{
		minute: CronField{fieldVals: []int{5}},
		hour: CronField{fieldVals: []int{0}},
		dayMonth: CronField{fieldVals: []int{Wildcard}},
		month: CronField{fieldVals: []int{8}},
		dayWeek: CronField{fieldVals: []int{Wildcard}},
	}

	assert.Equal(err, nil)

	assert.Equal(
		breakdown,
		result,
		"it returns a breakdown of the expression",
	)
}
