package base64_test

import (
	"testing"

	"github.com/bearchit/goboost/id"
	"github.com/bearchit/goboost/id/codec/base64"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	stdbase64 "encoding/base64"
)

func TestNewBase64URLCodec(t *testing.T) {
	coder := id.NewCoder(base64.NewURLCodec())

	t.Run("encoding", func(t *testing.T) {
		initialID := id.ID("home.featured")
		tid, err := coder.EncodeWith(initialID)
		require.NoError(t, err)

		decoded, err := stdbase64.URLEncoding.DecodeString(tid.String())
		require.NoError(t, err)
		assert.Equal(t, initialID, id.FromString(string(decoded)))
	})

	t.Run("encoding with no initial ID", func(t *testing.T) {
		_, err := coder.Encode()
		assert.Error(t, err)
	})

	t.Run("decoding", func(t *testing.T) {
		initialID := id.ID("home.featured")
		tid, err := coder.EncodeWith(initialID)
		require.NoError(t, err)

		decoded, payload, err := coder.Decode(tid)
		require.NoError(t, err)
		assert.True(t, payload.IsZero())
		assert.Equal(t, initialID, decoded)
	})
}
