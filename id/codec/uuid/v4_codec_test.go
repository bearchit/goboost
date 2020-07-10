package uuid_test

import (
	"testing"

	"github.com/bearchit/goboost/id/codec/uuid"

	"github.com/bearchit/goboost/id"

	libuuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewV4Codec(t *testing.T) {
	coder := id.NewCoder(uuid.NewV4Codec())

	t.Run("encoding", func(t *testing.T) {
		id, err := coder.Encode()
		require.NoError(t, err)

		parsed, err := libuuid.FromString(id.String())
		require.NoError(t, err)
		assert.Equal(t, libuuid.V4, parsed.Version())
	})

	t.Run("encoding with initial ID", func(t *testing.T) {
		_, err := coder.EncodeWith("initial")
		assert.Error(t, err)
	})

	t.Run("decoding", func(t *testing.T) {
		id, err := coder.Encode()
		require.NoError(t, err)

		decoded, payload, err := coder.Decode(id)
		require.NoError(t, err)
		assert.True(t, payload.IsZero())
		assert.Equal(t, id, decoded)
	})
}
