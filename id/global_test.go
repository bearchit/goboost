package id_test

import (
	"testing"

	"github.com/bearchit/goboost/id"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestNewEncoder(t *testing.T) {
	generator := id.NewGlobalGenerator("order")
	generated, err := generator.Generate()
	require.NoError(t, err)

	parser := id.NewGlobalParser()
	parsed, payload, err := parser.Parse(generated)
	require.NoError(t, err)

	parsedUUID, err := uuid.FromString(string(parsed))
	require.NoError(t, err)
	assert.Equal(t, uuid.V4, parsedUUID.Version())
	assert.Equal(t, "order", id.GlobalNamespace(payload))
}
