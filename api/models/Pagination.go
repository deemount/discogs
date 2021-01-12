package models

import (
	"net/url"
)

// PaginationRepository represents the contract
type PaginationRepository interface {
	Params() url.Values
}

// PaginationService ...
type PaginationService struct {
	Page      string
	PerPage   string
	Sort      string // year, title, format
	SortOrder string // asc, desc
}

// NewPaginationService is a constructor
func NewPaginationService(page, perpage, sort, order string) PaginationRepository {
	return &PaginationService{
		Page:      page,
		PerPage:   perpage,
		Sort:      sort,
		SortOrder: order,
	}
}

// Params converts pagination params to request values
func (p *PaginationService) Params() url.Values {

	if p == nil {
		return nil
	}

	params := url.Values{}

	params.Set("page", p.Page)
	params.Set("per_page", p.PerPage)
	params.Set("sort", p.Sort)
	params.Set("sort_order", p.SortOrder)

	return params

}
