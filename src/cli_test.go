package cron2e

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		args        []string
		expectedRes string
	}{
		{
			[]string{"cron2e", "5 0 * 3-8 *"},
			"Runs every day from months March through August at 00:05",
		},
	}

	for _, tc := range testCases {
		os.Args = tc.args
		os.Stdout = nil

		res := Run()

		assert.Equal(
			tc.expectedRes,
			res,
		)
	}
}
