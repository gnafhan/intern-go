package validation

type CreateTask struct {
	Title       string `json:"title" validate:"required,max=50" example:"fake title"`
	Description string `json:"description" validate:"required,max=255" example:"fake description"`
	CategoryID  string `json:"category_id" validate:"required,uuid" example:"fake category id"`
	Priority    string `json:"priority" validate:"required,oneof=low medium high" example:"low"`
	Deadline    string `json:"deadline" validate:"required,datetime=2006-01-02T15:04:05Z07:00" example:"2025-01-01T00:00:00Z"`
}

type UpdateTask struct {
	Title       string `json:"title,omitempty" validate:"omitempty,max=50" example:"fake title"`
	Description string `json:"description,omitempty" validate:"omitempty,max=255" example:"fake description"`
	CategoryID  string `json:"category_id,omitempty" validate:"omitempty,uuid" example:"fake category id"`
	Priority    string `json:"priority,omitempty" validate:"omitempty,oneof=low medium high" example:"medium"`
	Deadline    string `json:"deadline,omitempty" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00" example:"2025-01-01T00:00:00Z"`
}

type QueryTask struct {
	Page       int    `validate:"omitempty,number,max=50"`
	Limit      int    `validate:"omitempty,number,max=50"`
	Search     string `validate:"omitempty,max=50"`
	SortBy     string `validate:"omitempty,oneof=created_at updated_at deadline priority title"`
	SortOrder  string `validate:"omitempty,oneof=asc desc"`
	Priority   string `validate:"omitempty,oneof=low medium high"`
	CategoryID string `validate:"omitempty,uuid"`
}
