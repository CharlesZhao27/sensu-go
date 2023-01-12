package filter

import (
	"testing"

	corev3 "github.com/sensu/core/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoolean(t *testing.T) {
	fn := func(res corev3.Resource, val bool) bool {
		return false
	}

	f := Boolean(fn)
	require.NotNil(t, f)

	matches, err := f("true", nil)
	assert.NoError(t, err)
	assert.NotNil(t, matches)

	matches, err = f("false", nil)
	assert.NoError(t, err)
	assert.NotNil(t, matches)

	matches, err = f("  True ", nil)
	assert.NoError(t, err)
	assert.NotNil(t, matches)

	matches, err = f("truf ", nil)
	assert.Error(t, err)
	assert.Nil(t, matches)
}
