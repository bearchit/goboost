package id_test

import (
	"testing"

	"github.com/bearchit/goboost/id"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStubGenerator(t *testing.T) {
	g := id.NewStubGenerator("hello")
	generated, err := g.Generate()
	require.NoError(t, err)
	assert.Equal(t, id.ID("hello"), generated)
}
