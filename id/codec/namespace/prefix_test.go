package namespace_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/bearchit/goboost/id"

	"github.com/bearchit/goboost/id/codec/namespace"
)

func TestPrefixCodec_WithDelimiter(t *testing.T) {
	coder := id.NewCoder(namespace.NewPrefixCodec(
		"order",
		namespace.WithDelimiter("#"),
	))

	t.Run("encoding", func(t *testing.T) {
		initialID := id.ID("home.featured")
		tid, err := coder.EncodeWith(initialID)
		require.NoError(t, err)
		assert.Equal(t, id.FromString("order#home.featured"), tid)
	})

	t.Run("decoding", func(t *testing.T) {
		ns := "order"
		initialID := id.ID("home.featured")
		id, err := coder.EncodeWith(initialID)
		require.NoError(t, err)

		decoded, payload, err := coder.Decode(id)
		require.NoError(t, err)
		assert.Equal(t, initialID, decoded)
		assert.Equal(t, ns, payload[namespace.PayloadKeyNamespace])
	})
}

func TestPrefixCodec_DefaultDelimiter(t *testing.T) {
	coder := id.NewCoder(namespace.NewPrefixCodec(
		"order",
	))

	t.Run("encoding", func(t *testing.T) {
		initialID := id.ID("home.featured")
		tid, err := coder.EncodeWith(initialID)
		require.NoError(t, err)
		assert.Equal(t, id.FromString("order:home.featured"), tid)
	})

	t.Run("decoding", func(t *testing.T) {
		ns := "order"
		initialID := id.ID("home.featured")
		tid, err := coder.EncodeWith(initialID)
		require.NoError(t, err)

		decoded, payload, err := coder.Decode(tid)
		require.NoError(t, err)
		assert.Equal(t, initialID, decoded)
		assert.Equal(t, ns, payload[namespace.PayloadKeyNamespace])
	})
}

func TestDecodePrefix(t *testing.T) {
	t.Run("invalid prefix format", func(t *testing.T) {
		_, _, err := namespace.DecodePrefix(id.FromString("order"), ":")
		assert.Error(t, err)
	})
}
