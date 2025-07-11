package response

import "app/src/model"

type TaskResponse struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CategoryID  string         `json:"category_id"`
	Category    CategorySimple `json:"category"`
	Priority    string         `json:"priority"`
	Deadline    string         `json:"deadline"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
}

type CategorySimple struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TaskListResponse struct {
	Data    []TaskResponse `json:"data"`
	Total   int64          `json:"total"`
	Message string         `json:"message"`
}

type TaskDetailResponse struct {
	Data    TaskResponse `json:"data"`
	Message string       `json:"message"`
}

type TaskCreateResponse struct {
	Data    TaskResponse `json:"data"`
	Message string       `json:"message"`
}

type TaskUpdateResponse struct {
	Data    TaskResponse `json:"data"`
	Message string       `json:"message"`
}

type TaskDeleteResponse struct {
	Message string `json:"message"`
}

func ToTaskResponse(task *model.Task) TaskResponse {
	category := CategorySimple{}
	if task.Category.ID != (model.Category{}).ID {
		category.ID = task.Category.ID.String()
		category.Name = task.Category.Name
	} else {
		category.ID = task.CategoryID.String()
		category.Name = task.Category.Name
	}
	return TaskResponse{
		ID:          task.ID.String(),
		Title:       task.Title,
		Description: task.Description,
		CategoryID:  task.CategoryID.String(),
		Category:    category,
		Priority:    task.Priority,
		Deadline:    task.Deadline.Format("2006-01-02T15:04:05Z07:00"),
		CreatedAt:   task.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   task.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func ToTaskResponseList(tasks []model.Task) []TaskResponse {
	responses := make([]TaskResponse, len(tasks))
	for i, t := range tasks {
		responses[i] = ToTaskResponse(&t)
	}
	return responses
}
