package pagination_test

import (
	"testing"

	"github.com/bearchit/goboost/pagination"
	"github.com/stretchr/testify/assert"
)

func TestNewOffsetPageParam(t *testing.T) {
	t.Run("without default value", func(t *testing.T) {
		page, perPage := 1, 20
		param := pagination.NewOffsetPageParam(&page, &perPage)
		assert.Equal(t, page, param.Page)
		assert.Equal(t, perPage, param.PerPage)
	})

	t.Run("with default value", func(t *testing.T) {
		page, perPage := 2, 20
		param := pagination.NewOffsetPageParam(nil, nil,
			pagination.WithDefaultOffsetPageParam(page, perPage))
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
