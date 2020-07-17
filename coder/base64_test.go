package coder_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bearchit/goboost/coder"
)

func TestNewBase64URL(t *testing.T) {
	buf := new(bytes.Buffer)
	buf.WriteString("hello,world")
	value := buf.Bytes()

	encoder := coder.NewBase64URLEncoder()
	encoded, err := encoder.Encode(value)
	require.NoError(t, err)

	decoder := coder.NewBase64URLDecoder()
	decoded, payload, err := decoder.Decode(encoded)
	require.NoError(t, err)

	assert.Equal(t, value, decoded)
	assert.Nil(t, payload)
}
