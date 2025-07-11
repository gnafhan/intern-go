package response

type CategoryResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CategoryListResponse struct {
	Data    []CategoryResponse `json:"data"`
	Total   int64              `json:"total"`
	Message string             `json:"message"`
}

type CategoryDetailResponse struct {
	Data    CategoryResponse `json:"data"`
	Message string           `json:"message"`
}

type CategoryCreateResponse struct {
	Data    CategoryResponse `json:"data"`
	Message string           `json:"message"`
}

type CategoryUpdateResponse struct {
	Data    CategoryResponse `json:"data"`
	Message string           `json:"message"`
}

type CategoryDeleteResponse struct {
	Message string `json:"message"`
}
