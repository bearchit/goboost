package coder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bearchit/goboost/coder"
)

func TestNewPrefixEncoder(t *testing.T) {
	t.Run("without option", func(t *testing.T) {
		encoder := coder.NewPrefixEncoder([]byte("test"))
		encoded, err := encoder.Encode([]byte("hello"))
		require.NoError(t, err)
		assert.Equal(t, []byte("test:hello"), encoded)
	})

	t.Run("with delimiter", func(t *testing.T) {
		encoder := coder.NewPrefixEncoder([]byte("test"), coder.PrefixDelim([]byte("#")))
		encoded, err := encoder.Encode([]byte("hello"))
		require.NoError(t, err)
		assert.Equal(t, []byte("test#hello"), encoded)
	})
}

func TestNewPrefixDecoder(t *testing.T) {
	t.Run("without option", func(t *testing.T) {
		decoder := coder.NewPrefixDecoder()
		decoded, payload, err := decoder.Decode([]byte("test:hello"))
		require.NoError(t, err)
		assert.Equal(t, []byte("hello"), decoded)

		v, err := payload.Get(coder.PayloadPrefix)
		require.NoError(t, err)
		assert.Equal(t, []byte("test"), v)
	})

	t.Run("with delimiter", func(t *testing.T) {
		decoder := coder.NewPrefixDecoder(coder.PrefixDelim([]byte("#")))
		decoded, payload, err := decoder.Decode([]byte("test#hello"))
		require.NoError(t, err)
		assert.Equal(t, []byte("hello"), decoded)

		v, err := payload.Get(coder.PayloadPrefix)
		require.NoError(t, err)
		assert.Equal(t, []byte("test"), v)
	})
}
