package api

import "github.com/lucasti79/meli-interview/internal/category"

// swagger:model CategoryResult
type CategoryResult struct {
	Data []category.Category `json:"data"`
}

// swagger:model CategoriesResult
type CategoriesResult struct {
	Data *[]category.Category `json:"data"`
}
