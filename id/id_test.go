package id_test

import (
	"testing"

	"github.com/bearchit/goboost/id"
	"github.com/stretchr/testify/assert"
)

func TestNewIDs(t *testing.T) {
	ids := id.NewIDs("1", "2")
	assert.Equal(t, id.IDs{"1", "2"}, ids)
}
