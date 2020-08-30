package pagination

import (
	"math"

	"github.com/bearchit/goboost/nullable"
)

type CursorPageParam struct {
	First  *int
	Last   *int
	Before *string
	After  *string
}

func NewCursorPageParam(first, last *int, before, after *string) CursorPageParam {
	return CursorPageParam{
		First:  first,
		Last:   last,
		Before: before,
		After:  after,
	}
}

type CursorPageInfo struct {
	StartCursor     string
	EndCursor       string
	HasPreviousPage bool
	HasNextPage     bool
}

type OffsetPageParam struct {
	Page    int
	PerPage int
}

func (p OffsetPageParam) Offset() int64 {
	return int64((p.Page - 1) * p.PerPage)
}

func NewOffsetPageParam(page, perPage *int, option ...OffsetPageParamOptionFunc) OffsetPageParam {
	param := OffsetPageParam{
		Page:    nullable.PtrToInt(page),
		PerPage: nullable.PtrToInt(perPage),
	}

	for _, x := range option {
		x(&param)
	}

	return param
}

type OffsetPageParamOptionFunc func(*OffsetPageParam)

func WithDefaultOffsetPageParam(page, perPage int) OffsetPageParamOptionFunc {
	return func(param *OffsetPageParam) {
		param.Page = page
		param.PerPage = perPage
	}
}

type OffsetPageInfo struct {
	Page            int
	PerPage         int
	TotalPages      int
	HasPreviousPage bool
	HasNextPage     bool
}

func NewOffsetPageInfoFrom(
	page int,
	perPage int,
	totalItems int,
) OffsetPageInfo {
	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))
	return OffsetPageInfo{
		Page:            page,
		PerPage:         perPage,
		TotalPages:      totalPages,
		HasPreviousPage: page > 1,
		HasNextPage:     page < totalPages,
	}
}
