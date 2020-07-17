package coder_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/bearchit/goboost/coder"
)

func TestNewUuidV4Encoder(t *testing.T) {
	encoder := coder.NewUuidV4Encoder()
	encoded, err := encoder.Encode(nil)
	require.NoError(t, err)

	parsed, err := uuid.FromString(string(encoded))
	require.NoError(t, err)
	assert.Equal(t, uuid.V4, parsed.Version())
}
