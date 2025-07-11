package service

import (
	"app/src/model"
	"app/src/utils"
	"app/src/validation"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TaskService interface {
	GetTasks(c *fiber.Ctx, params *validation.QueryTask) ([]model.Task, int64, error)
	GetTaskByID(c *fiber.Ctx, id string) (*model.Task, error)
	CreateTask(c *fiber.Ctx, req *validation.CreateTask) (*model.Task, error)
	UpdateTask(c *fiber.Ctx, req *validation.UpdateTask, id string) (*model.Task, error)
	DeleteTask(c *fiber.Ctx, id string) error
}

type taskService struct {
	Log      *logrus.Logger
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewTaskService(db *gorm.DB, validate *validator.Validate) TaskService {
	return &taskService{
		Log:      utils.Log,
		DB:       db,
		Validate: validate,
	}
}

func (s *taskService) GetTasks(c *fiber.Ctx, params *validation.QueryTask) ([]model.Task, int64, error) {
	var tasks []model.Task
	var total int64

	query := s.DB.Model(&model.Task{}).Preload("Category")

	if params.Search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+params.Search+"%", "%"+params.Search+"%")
	}
	if params.Priority != "" {
		query = query.Where("priority = ?", params.Priority)
	}
	if params.CategoryID != "" {
		query = query.Where("category_id = ?", params.CategoryID)
	}

	if err := query.Count(&total).Error; err != nil {
		s.Log.Errorf("Failed to count tasks: %v", err)
		return nil, 0, err
	}

	// Sorting
	sortBy := params.SortBy
	if sortBy == "" {
		sortBy = "created_at"
	}
	sortOrder := params.SortOrder
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}
	// Only allow certain fields to be sorted
	allowedSortFields := map[string]bool{"created_at": true, "updated_at": true, "deadline": true, "priority": true, "title": true}
	if !allowedSortFields[sortBy] {
		sortBy = "created_at"
	}
	query = query.Order(sortBy + " " + sortOrder)

	query = query.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)

	if err := query.Find(&tasks).Error; err != nil {
		s.Log.Errorf("Failed to get tasks: %v", err)
		return nil, 0, err
	}

	return tasks, total, nil
}

func (s *taskService) GetTaskByID(c *fiber.Ctx, id string) (*model.Task, error) {
	task := new(model.Task)

	result := s.DB.WithContext(c.Context()).Preload("Category").First(task, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fiber.NewError(fiber.StatusNotFound, "Task not found")
	}

	return task, nil
}

func (s *taskService) CreateTask(c *fiber.Ctx, req *validation.CreateTask) (*model.Task, error) {
	var task model.Task

	categoryID, err := utils.ParseUUID(req.CategoryID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid category_id")
	}

	deadline, err := time.Parse(time.RFC3339, req.Deadline)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid deadline format")
	}

	task.Title = req.Title
	task.Description = req.Description
	task.CategoryID = categoryID
	task.Priority = req.Priority
	task.Deadline = deadline

	if err := s.DB.Create(&task).Error; err != nil {
		s.Log.Errorf("Failed to create task: %v", err)
		return nil, err
	}

	// Preload category for response
	if err := s.DB.Preload("Category").First(&task, "id = ?", task.ID).Error; err != nil {
		s.Log.Errorf("Failed to preload category after create: %v", err)
		return nil, err
	}

	return &task, nil
}

func (s *taskService) UpdateTask(c *fiber.Ctx, req *validation.UpdateTask, id string) (*model.Task, error) {
	// Update using a map, then reload from DB to get the latest state
	updates := map[string]interface{}{}

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.CategoryID != "" {
		categoryID, err := utils.ParseUUID(req.CategoryID)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid category_id")
		}
		updates["category_id"] = categoryID
	}
	if req.Priority != "" {
		updates["priority"] = req.Priority
	}
	if req.Deadline != "" {
		deadline, err := time.Parse(time.RFC3339, req.Deadline)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid deadline format")
		}
		updates["deadline"] = deadline
	}

	if len(updates) == 0 {
		return s.GetTaskByID(c, id) // nothing to update, just return current
	}

	result := s.DB.Model(&model.Task{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		s.Log.Errorf("Failed to update task: %v", result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Task not found")
	}

	var updatedTask model.Task
	if err := s.DB.Preload("Category").First(&updatedTask, "id = ?", id).Error; err != nil {
		s.Log.Errorf("Failed to preload category after update: %v", err)
		return nil, err
	}

	return &updatedTask, nil
}

func (s *taskService) DeleteTask(c *fiber.Ctx, id string) error {
	_, err := s.GetTaskByID(c, id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Task not found")
	}

	if err := s.DB.Delete(&model.Task{}, "id = ?", id).Error; err != nil {
		s.Log.Errorf("Failed to delete task: %v", err)
		return err
	}

	return nil
}
