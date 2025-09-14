package api

import "github.com/lucasti79/meli-interview/internal/product"

// swagger:model ProductPaginatedResult
type ProductPaginatedResult struct {
	Data       []product.Product `json:"data"`
	TotalCount int               `json:"totalCount"`
	Page       int               `json:"page"`
	PageSize   int               `json:"pageSize"`
}

// swagger:model ProductResult
type ProductResult struct {
	Data *product.Product `json:"data"`
}
