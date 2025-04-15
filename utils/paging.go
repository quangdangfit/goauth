package utils

import commonPb "github.com/quangdangfit/goauth/proto/api/common"

const DefaultLimit = 100
const MaxLimit = 5000

type Pagination struct {
	limit  int32
	offset int32
	page   int32
	total  int32
}

func NewPagination(page, limit int) *Pagination {
	if limit < 1 {
		limit = DefaultLimit
	}
	if limit > MaxLimit {
		limit = MaxLimit
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	return &Pagination{
		offset: int32(offset),
		page:   int32(page),
		limit:  int32(limit),
	}
}

func (p *Pagination) Limit() int {
	return int(p.limit)
}

func (p *Pagination) Offset() int {
	return int(p.offset)
}

func (p *Pagination) SetTotal(total int32) {
	p.total = total
}

func (p *Pagination) Proto() *commonPb.PaginationResponse {
	return &commonPb.PaginationResponse{
		Total:  p.total,
		Page:   p.page,
		Limit:  p.limit,
		Offset: p.offset,
	}
}
