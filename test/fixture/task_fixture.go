package fixture

import (
	"app/src/model"
	"time"

	"github.com/google/uuid"
)

var TaskOne = &model.Task{
	ID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
	Title:       "Task One",
	Description: "Description for Task One",
	CategoryID:  uuid.MustParse("22222222-2222-2222-2222-222222222222"),
	Priority:    "high",
	Deadline:    time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
}

var TaskTwo = &model.Task{
	ID:          uuid.MustParse("33333333-3333-3333-3333-333333333333"),
	Title:       "Task Two",
	Description: "Description for Task Two",
	CategoryID:  uuid.MustParse("22222222-2222-2222-2222-222222222222"),
	Priority:    "medium",
	Deadline:    time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC),
}
