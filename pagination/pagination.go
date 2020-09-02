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

func (p OffsetPageParam) PageInfo(totalItems int) OffsetPageInfo {
	return NewOffsetPageInfoFrom(p.Page, p.PerPage, totalItems)
}

func NewOffsetPageParam(page, perPage int) OffsetPageParam {
	return OffsetPageParam{
		Page:    page,
		PerPage: perPage,
	}
}

type OffsetPageParamFactory struct {
	page    int
	perPage int
}

func NewOffsetPageParamFactory(page, perPage int) OffsetPageParamFactory {
	return OffsetPageParamFactory{
		page:    page,
		perPage: perPage,
	}
}

func (f OffsetPageParamFactory) Param(page, perPage *int) OffsetPageParam {
	return OffsetPageParam{
		Page:    nullable.DefaultIfNilInt(page, f.page),
		PerPage: nullable.DefaultIfNilInt(perPage, f.perPage),
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
