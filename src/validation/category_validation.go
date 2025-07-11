package validation

type CreateCategory struct {
	Name string `json:"name" validate:"required,max=50" example:"fake name"`
}

type UpdateCategory struct {
	Name string `json:"name" validate:"required,max=50" example:"fake name"`
}

type QueryCategory struct {
	Page   int    `validate:"omitempty,number,max=50"`
	Limit  int    `validate:"omitempty,number,max=50"`
	Search string `validate:"omitempty,max=50"`
}
