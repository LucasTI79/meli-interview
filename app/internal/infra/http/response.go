package httpdto

type PaginatedResult[T any] struct {
	Data       []T `json:"data"`
	TotalCount int `json:"totalCount"`
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Result[T any] struct {
	Data T `json:"data,omitempty"`
}
