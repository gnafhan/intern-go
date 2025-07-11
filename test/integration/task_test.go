package integration

import (
	"app/src/model"
	"app/src/validation"
	"app/test"
	"app/test/fixture"
	"app/test/helper"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTaskRoutes(t *testing.T) {
	t.Run("POST /v1/tasks", func(t *testing.T) {
		t.Run("should return 201 and create new task if data is ok", func(t *testing.T) {
			helper.ClearAll(test.DB)
			helper.ClearTasks(test.DB)
			helper.ClearCategories(test.DB)
			// Insert required category
			category := &model.Category{ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"), Name: "Category One"}
			test.DB.Create(category)

			newTask := validation.CreateTask{
				Title:       "New Task",
				Description: "A new task description",
				CategoryID:  category.ID.String(),
				Priority:    "high",
				Deadline:    time.Now().Add(24 * time.Hour).UTC().Format(time.RFC3339),
			}

			bodyJSON, _ := json.Marshal(newTask)
			request := httptest.NewRequest(http.MethodPost, "/v1/tasks", strings.NewReader(string(bodyJSON)))
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("Accept", "application/json")

			apiResponse, err := test.App.Test(request)
			assert.Nil(t, err)
			assert.Equal(t, http.StatusCreated, apiResponse.StatusCode)

			bytes, _ := io.ReadAll(apiResponse.Body)
			var resp map[string]interface{}
			_ = json.Unmarshal(bytes, &resp)
			assert.Equal(t, "Task created successfully", resp["message"])
		})
	})

	t.Run("GET /v1/tasks", func(t *testing.T) {
		helper.ClearAll(test.DB)
		helper.ClearTasks(test.DB)
		helper.ClearCategories(test.DB)
		category := &model.Category{ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"), Name: "Category One"}
		test.DB.Create(category)
		task1 := *fixture.TaskOne
		test.DB.Create(&task1)
		task2 := *fixture.TaskTwo
		test.DB.Create(&task2)

		request := httptest.NewRequest(http.MethodGet, "/v1/tasks?page=1&limit=10", nil)
		apiResponse, err := test.App.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, apiResponse.StatusCode)

		bytes, _ := io.ReadAll(apiResponse.Body)
		var resp map[string]interface{}
		_ = json.Unmarshal(bytes, &resp)
		assert.Equal(t, "Tasks retrieved successfully", resp["message"])
		assert.NotNil(t, resp["data"])
	})

	t.Run("GET /v1/tasks/:taskId", func(t *testing.T) {
		helper.ClearAll(test.DB)
		helper.ClearTasks(test.DB)
		helper.ClearCategories(test.DB)
		category := &model.Category{ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"), Name: "Category One"}
		test.DB.Create(category)
		task3 := *fixture.TaskOne
		test.DB.Create(&task3)

		// Debug: check if the task exists in the DB
		var task model.Task
		err := test.DB.First(&task, "id = ?", fixture.TaskOne.ID).Error
		assert.Nil(t, err, "Task should exist in DB after creation")

		request := httptest.NewRequest(http.MethodGet, "/v1/tasks/"+fixture.TaskOne.ID.String(), nil)
		apiResponse, err := test.App.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, apiResponse.StatusCode)

		bytes, _ := io.ReadAll(apiResponse.Body)
		var resp map[string]interface{}
		_ = json.Unmarshal(bytes, &resp)
		assert.Equal(t, "Task retrieved successfully", resp["message"])
		assert.NotNil(t, resp["data"])
	})

	t.Run("PUT /v1/tasks/:taskId", func(t *testing.T) {
		helper.ClearAll(test.DB)
		helper.ClearTasks(test.DB)
		helper.ClearCategories(test.DB)
		category := &model.Category{ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"), Name: "Category One"}
		test.DB.Create(category)
		task4 := *fixture.TaskOne
		test.DB.Create(&task4)

		updateTask := validation.UpdateTask{
			Title: "Updated Task Title",
		}
		bodyJSON, _ := json.Marshal(updateTask)
		request := httptest.NewRequest(http.MethodPut, "/v1/tasks/"+fixture.TaskOne.ID.String(), strings.NewReader(string(bodyJSON)))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept", "application/json")

		apiResponse, err := test.App.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, apiResponse.StatusCode)

		bytes, _ := io.ReadAll(apiResponse.Body)
		var resp map[string]interface{}
		_ = json.Unmarshal(bytes, &resp)
		assert.Equal(t, "Task updated successfully", resp["message"])
	})

	t.Run("DELETE /v1/tasks/:taskId", func(t *testing.T) {
		helper.ClearAll(test.DB)
		helper.ClearTasks(test.DB)
		helper.ClearCategories(test.DB)
		category := &model.Category{ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"), Name: "Category One"}
		test.DB.Create(category)
		task5 := *fixture.TaskOne
		test.DB.Create(&task5)

		request := httptest.NewRequest(http.MethodDelete, "/v1/tasks/"+fixture.TaskOne.ID.String(), nil)
		apiResponse, err := test.App.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, apiResponse.StatusCode)

		bytes, _ := io.ReadAll(apiResponse.Body)
		var resp map[string]interface{}
		_ = json.Unmarshal(bytes, &resp)
		assert.Equal(t, "Task deleted successfully", resp["message"])
	})
}
