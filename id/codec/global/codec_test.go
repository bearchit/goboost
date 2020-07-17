package global_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bearchit/goboost/id/codec/global"
)

func TestNewEncoder(t *testing.T) {
	generator := global.NewGenerator("order")
	generated, err := generator.Generate()
	require.NoError(t, err)
	t.Log(generated)

	parser := global.NewParser()
	parsed, payload, err := parser.Parse(generated)
	require.NoError(t, err)
	t.Log(parsed, global.Namespace(payload))
}
