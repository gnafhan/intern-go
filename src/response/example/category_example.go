package example

type CategoryExample struct {
	ID        string `json:"id" example:"fake id"`
	Name      string `json:"name" example:"fake name"`
	CreatedAt string `json:"created_at" example:"2025-01-01T00:00:00Z"`
	UpdatedAt string `json:"updated_at" example:"2025-01-01T00:00:00Z"`
}

type GetCategoriesResponse struct {
	Message string            `json:"message" example:"Categories retrieved successfully"`
	Data    []CategoryExample `json:"data"`
	Total   int64             `json:"total" example:"100"`
	Status  string            `json:"status" example:"success"`
}

type GetCategoryByIDResponse struct {
	Message string          `json:"message" example:"Category retrieved successfully"`
	Data    CategoryExample `json:"data"`
	Status  string          `json:"status" example:"success"`
}

type CreateCategoryResponse struct {
	Message string          `json:"message" example:"Category created successfully"`
	Data    CategoryExample `json:"data"`
	Status  string          `json:"status" example:"success"`
}

type UpdateCategoryResponse struct {
	Message string          `json:"message" example:"Category updated successfully"`
	Data    CategoryExample `json:"data"`
	Status  string          `json:"status" example:"success"`
}

type DeleteCategoryResponse struct {
	Message string `json:"message" example:"Category deleted successfully"`
	Status  string `json:"status" example:"success"`
}
