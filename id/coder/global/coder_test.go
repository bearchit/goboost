package global_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"

	"github.com/bearchit/goboost/id"
	"github.com/bearchit/goboost/id/codec/namespace"
	"github.com/stretchr/testify/assert"

	"github.com/bearchit/goboost/id/coder/global"
)

func TestNewCoder(t *testing.T) {
	coder := global.NewCoder("order", global.WithDelimiter("#"))
	gid := coder.MustEncode()

	decoded, payload := coder.MustDecode(gid)
	parsedUUID, err := uuid.FromString(decoded.String())
	require.NoError(t, err)
	assert.Equal(t, decoded, id.FromString(parsedUUID.String()))
	assert.Equal(t, "order", payload[namespace.PayloadKeyNamespace])
}
