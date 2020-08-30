package structs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bearchit/goboost/structs"
	"github.com/stretchr/testify/require"
)

func TestFieldByTag(t *testing.T) {
	s := struct {
		ID   int `crud:"pk"`
		Name string
	}{
		ID:   1,
		Name: "Tester",
	}

	t.Run("field by tag keys, name pairs", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			f, err := structs.FieldByTagKeyNamePairs(s, []string{"crud", "pk"})
			require.NoError(t, err)
			assert.Equal(t, "ID", f.Name)
			assert.Equal(t, 1, f.Value)
		})

		t.Run("fail", func(t *testing.T) {
			f, err := structs.FieldByTagKeyNamePairs(s, []string{"crud", "name"})
			assert.Error(t, err)
			assert.Nil(t, f)
		})
	})
}
