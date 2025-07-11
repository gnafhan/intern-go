package example

type TaskExample struct {
	ID          string           `json:"id" example:"fake id"`
	Title       string           `json:"title" example:"fake title"`
	Description string           `json:"description" example:"fake description"`
	CategoryID  string           `json:"category_id" example:"fake category id"`
	Category    CategorySimpleEx `json:"category"`
	Priority    string           `json:"priority" example:"high"`
	Deadline    string           `json:"deadline" example:"2025-01-01T00:00:00Z"`
	CreatedAt   string           `json:"created_at" example:"2025-01-01T00:00:00Z"`
	UpdatedAt   string           `json:"updated_at" example:"2025-01-01T00:00:00Z"`
}

type CategorySimpleEx struct {
	ID   string `json:"id" example:"fake category id"`
	Name string `json:"name" example:"fake category name"`
}

type GetTasksResponse struct {
	Message string        `json:"message" example:"Tasks retrieved successfully"`
	Data    []TaskExample `json:"data"`
	Total   int64         `json:"total" example:"100"`
	Status  string        `json:"status" example:"success"`
}

type GetTaskByIDResponse struct {
	Message string      `json:"message" example:"Task retrieved successfully"`
	Data    TaskExample `json:"data"`
	Status  string      `json:"status" example:"success"`
}

type CreateTaskResponse struct {
	Message string      `json:"message" example:"Task created successfully"`
	Data    TaskExample `json:"data"`
	Status  string      `json:"status" example:"success"`
}

type UpdateTaskResponse struct {
	Message string      `json:"message" example:"Task updated successfully"`
	Data    TaskExample `json:"data"`
	Status  string      `json:"status" example:"success"`
}

type DeleteTaskResponse struct {
	Message string `json:"message" example:"Task deleted successfully"`
	Status  string `json:"status" example:"success"`
}
