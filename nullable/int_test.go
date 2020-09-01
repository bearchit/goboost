package nullable_test

import (
	"testing"

	"github.com/bearchit/goboost/nullable"
	"github.com/stretchr/testify/assert"
)

func TestDefaultIfNilInt(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		v := nullable.DefaultIfNilInt(nil, 5)
		assert.Equal(t, 5, v)
	})

	t.Run("value", func(t *testing.T) {
		v := nullable.DefaultIfNilInt(nullable.IntToPtr(10), 15)
		assert.Equal(t, 10, v)
	})
}
