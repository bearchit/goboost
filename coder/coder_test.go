package coder_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/bearchit/goboost/coder"
)

func TestPayload(t *testing.T) {
	payload := coder.NewPayload()
	err := payload.Set("msg", "hello, world")
	require.NoError(t, err)

	v, err := payload.Get("msg")
	require.NoError(t, err)
	assert.Equal(t, "hello, world", v.(string))

	t.Run("key already exists", func(t *testing.T) {
		var err error
		payload := coder.NewPayload()
		err = payload.Set("msg", "hello, world")
		require.NoError(t, err)
		err = payload.Set("msg", "hello, goboost")
		assert.True(t, errors.Is(err, coder.ErrPayloadExists))
	})

	t.Run("key not found", func(t *testing.T) {
		payload := coder.NewPayload()
		_, err := payload.Get("msg")
		assert.True(t, errors.Is(err, coder.ErrPayloadNotFound))
	})
}
