package id_test

import (
	"testing"

	"github.com/bearchit/goboost/id"
	"github.com/stretchr/testify/assert"
)

func TestNewIDs(t *testing.T) {
	ids := id.NewIDs(id.ID("1"), id.ID("2"))
	assert.Equal(t, id.IDs{id.ID("1"), id.ID("2")}, ids)
}
