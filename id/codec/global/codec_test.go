package global_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/bearchit/goboost/id/codec/global"
)

func TestNewEncoder(t *testing.T) {
	generator := global.NewGenerator("order")
	generated, err := generator.Generate()
	require.NoError(t, err)

	parser := global.NewParser()
	parsed, payload, err := parser.Parse(generated)
	require.NoError(t, err)

	parsedUUID, err := uuid.FromString(string(parsed))
	require.NoError(t, err)
	assert.Equal(t, uuid.V4, parsedUUID.Version())
	assert.Equal(t, "order", global.Namespace(payload))
}
