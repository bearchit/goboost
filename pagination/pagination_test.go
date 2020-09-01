package pagination_test

import (
	"testing"

	"github.com/bearchit/goboost/nullable"

	"github.com/bearchit/goboost/pagination"
	"github.com/stretchr/testify/assert"
)

func TestNewOffsetPageParam(t *testing.T) {
	t.Run("without default value", func(t *testing.T) {
		page, perPage := 1, 20
		param := pagination.NewOffsetPageParam(page, perPage)
		assert.Equal(t, page, param.Page)
		assert.Equal(t, perPage, param.PerPage)
	})
}

func TestNewOffsetPageInfoFrom(t *testing.T) {
	pageInfo := pagination.NewOffsetPageInfoFrom(1, 20, 100)
	assert.Equal(t, 1, pageInfo.Page)
	assert.Equal(t, 20, pageInfo.PerPage)
	assert.Equal(t, 5, pageInfo.TotalPages)
	assert.False(t, pageInfo.HasPreviousPage)
	assert.True(t, pageInfo.HasNextPage)
}

func TestOffsetPageParamFactory(t *testing.T) {
	factory := pagination.NewOffsetPageParamFactory(1, 20)

	t.Run("default", func(t *testing.T) {
		param := factory.Param(nil, nil)

		assert.Equal(t, 1, param.Page)
		assert.Equal(t, 20, param.PerPage)
	})

	t.Run("value", func(t *testing.T) {
		param := factory.Param(nullable.IntToPtr(1), nullable.IntToPtr(10))

		assert.Equal(t, 1, param.Page)
		assert.Equal(t, 10, param.PerPage)
	})
}
