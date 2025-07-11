package model_test

import (
	"app/src/model"
	"app/src/validation"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskModel(t *testing.T) {
	t.Run("Create task validation", func(t *testing.T) {
		var newTask = validation.CreateTask{
			Title:       "Test Task",
			Description: "A test task description",
			CategoryID:  "22222222-2222-2222-2222-222222222222",
			Priority:    "high",
			Deadline:    "2025-01-01T00:00:00Z",
		}

		t.Run("should correctly validate a valid task", func(t *testing.T) {
			err := validation.Validator().Struct(newTask)
			assert.NoError(t, err)
		})

		t.Run("should throw a validation error if title is missing", func(t *testing.T) {
			newTask.Title = ""
			err := validation.Validator().Struct(newTask)
			assert.Error(t, err)
			newTask.Title = "Test Task" // reset
		})

		t.Run("should throw a validation error if description is missing", func(t *testing.T) {
			newTask.Description = ""
			err := validation.Validator().Struct(newTask)
			assert.Error(t, err)
			newTask.Description = "A test task description"
		})

		t.Run("should throw a validation error if category_id is invalid", func(t *testing.T) {
			newTask.CategoryID = "not-a-uuid"
			err := validation.Validator().Struct(newTask)
			assert.Error(t, err)
			newTask.CategoryID = "22222222-2222-2222-2222-222222222222"
		})

		t.Run("should throw a validation error if priority is invalid", func(t *testing.T) {
			newTask.Priority = "urgent"
			err := validation.Validator().Struct(newTask)
			assert.Error(t, err)
			newTask.Priority = "high"
		})

		t.Run("should throw a validation error if deadline is invalid format", func(t *testing.T) {
			newTask.Deadline = "not-a-date"
			err := validation.Validator().Struct(newTask)
			assert.Error(t, err)
			newTask.Deadline = "2025-01-01T00:00:00Z"
		})
	})

	t.Run("Update task validation", func(t *testing.T) {
		var updateTask = validation.UpdateTask{
			Title:       "Updated Task",
			Description: "Updated description",
			CategoryID:  "22222222-2222-2222-2222-222222222222",
			Priority:    "medium",
			Deadline:    "2025-06-01T00:00:00Z",
		}

		t.Run("should correctly validate a valid update", func(t *testing.T) {
			err := validation.Validator().Struct(updateTask)
			assert.NoError(t, err)
		})

		t.Run("should throw a validation error if category_id is invalid", func(t *testing.T) {
			updateTask.CategoryID = "not-a-uuid"
			err := validation.Validator().Struct(updateTask)
			assert.Error(t, err)
			updateTask.CategoryID = "22222222-2222-2222-2222-222222222222"
		})

		t.Run("should throw a validation error if priority is invalid", func(t *testing.T) {
			updateTask.Priority = "urgent"
			err := validation.Validator().Struct(updateTask)
			assert.Error(t, err)
			updateTask.Priority = "medium"
		})

		t.Run("should throw a validation error if deadline is invalid format", func(t *testing.T) {
			updateTask.Deadline = "not-a-date"
			err := validation.Validator().Struct(updateTask)
			assert.Error(t, err)
			updateTask.Deadline = "2025-06-01T00:00:00Z"
		})
	})

	t.Run("Task toJSON()", func(t *testing.T) {
		t.Run("should marshal task to JSON with all fields", func(t *testing.T) {
			task := &model.Task{
				Title:       "Test Task",
				Description: "A test task description",
				Priority:    "high",
			}
			bytes, _ := json.Marshal(task)
			assert.Contains(t, string(bytes), "Test Task")
			assert.Contains(t, string(bytes), "A test task description")
			assert.Contains(t, string(bytes), "high")
		})
	})
}
