package awsrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	assert := assert.New(t)

	result, errs := format.Translate("rate(8 hours)")

	assert.Equal(0, len(errs))
	assert.Equal("Runs every 8 hours", result)
}
