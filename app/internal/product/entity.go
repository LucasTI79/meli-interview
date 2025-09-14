package product

type Product struct {
	Id            string  `json:"productId"`
	Description   string  `json:"description"`
	Name          string  `json:"name"`
	OriginalPrice float64 `json:"originalPrice"`
	Price         float64 `json:"price"`
	Category      string  `json:"category"`
	Image         string  `json:"image"`
	InStock       bool    `json:"inStock"`
	Rating        float64 `json:"rating"`
	Reviews       int     `json:"reviews"`
}

// swagger:parameters GetAll
type ProductFilter struct {
	// in: query
	Name string `json:"name,omitempty" validate:"omitempty"`
	// in: query
	Categories []string `json:"categories,omitempty" validate:"omitempty"`
	// in: query
	MinPrice float64 `json:"minPrice,omitempty" validate:"omitempty"`
	// in: query
	MaxPrice float64 `json:"maxPrice,omitempty" validate:"omitempty"`
	// in: query
	Page int `json:"page,omitempty" validate:"omitempty,min=1"`
	// in: query
	PageSize int `json:"pageSize,omitempty" validate:"omitempty,min=1,max=100"`
}
