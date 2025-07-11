package controller

import (
	"app/src/response"
	"app/src/service"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	TaskService service.TaskService
}

func NewTaskController(taskService service.TaskService) *TaskController {
	return &TaskController{
		TaskService: taskService,
	}
}

// @Tags         Tasks
// @Summary      Get all tasks
// @Description  Retrieve all tasks.
// @Produce      json
// @Param        page     query     int     false   "Page number"  default(0)
// @Param        limit    query     int     false   "Maximum number of tasks"    default(10)
// @Param        search   query     string  false  "Search by title or description"
// @Param        sort_by   query     string  false  "Sort by field (created_at, updated_at, deadline, priority, title)"
// @Param        sort_order query    string  false  "Sort order (asc, desc)"
// @Param        priority  query     string  false  "Filter by priority (low, medium, high)"
// @Param        category_id query   string  false  "Filter by category_id (UUID)"
// @Success      200  {object}  response.TaskListResponse
// @Failure      500  {object}  response.ErrorResponse  "Internal server error"
// @Router       /tasks [get]
func (c *TaskController) GetTasks(ctx *fiber.Ctx) error {
	query := &validation.QueryTask{
		Page:       ctx.QueryInt("page", 1),
		Limit:      ctx.QueryInt("limit", 10),
		Search:     ctx.Query("search", ""),
		SortBy:     ctx.Query("sort_by", ""),
		SortOrder:  ctx.Query("sort_order", ""),
		Priority:   ctx.Query("priority", ""),
		CategoryID: ctx.Query("category_id", ""),
	}
	tasks, total, err := c.TaskService.GetTasks(ctx, query)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get tasks",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    response.ToTaskResponseList(tasks),
		"total":   total,
		"message": "Tasks retrieved successfully",
		"status":  "success",
	})
}

// @Tags         Tasks
// @Summary      Get a task by ID
// @Description  Retrieve a task by ID.
// @Produce      json
// @Param        taskId   path     string  true  "Task ID"
// @Success      200  {object}  response.TaskDetailResponse
// @Failure      404  {object}  response.ErrorResponse  "Task not found"
// @Router       /tasks/{taskId} [get]
func (c *TaskController) GetTaskByID(ctx *fiber.Ctx) error {
	taskId := ctx.Params("taskId")
	task, err := c.TaskService.GetTaskByID(ctx, taskId)
	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok && fiberErr.Code == fiber.StatusNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Task not found",
				"status":  "error",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get task",
			"status":  "error",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    response.ToTaskResponse(task),
		"message": "Task retrieved successfully",
		"status":  "success",
	})
}

// @Tags         Tasks
// @Summary      Create a task
// @Description  Create a task.
// @Produce      json
// @Param        task   body     validation.CreateTask  true  "Task"
// @Success      201  {object}  response.TaskCreateResponse
// @Failure      400  {object}  response.ErrorResponse  "Invalid request body"
// @Failure      500  {object}  response.ErrorResponse  "Internal server error"
// @Router       /tasks [post]
func (c *TaskController) CreateTask(ctx *fiber.Ctx) error {
	req := new(validation.CreateTask)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  "error",
		})
	}
	validate := validation.Validator()
	if err := validate.Struct(req); err != nil {
		errorsMap := validation.CustomErrorMessages(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  "error",
			"errors":  errorsMap,
		})
	}
	task, err := c.TaskService.CreateTask(ctx, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create task",
			"status":  "error",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    response.ToTaskResponse(task),
		"message": "Task created successfully",
		"status":  "success",
	})
}

// @Tags         Tasks
// @Summary      Update a task
// @Description  Update a task.
// @Produce      json
// @Param        taskId   path     string  true  "Task ID"
// @Param        task   body     validation.UpdateTask  true  "Task"
// @Success      200  {object}  response.TaskUpdateResponse
// @Failure      400  {object}  response.ErrorResponse  "Invalid request body"
// @Failure      500  {object}  response.ErrorResponse  "Internal server error"
// @Router       /tasks/{taskId} [put]
func (c *TaskController) UpdateTask(ctx *fiber.Ctx) error {
	taskId := ctx.Params("taskId")
	req := new(validation.UpdateTask)
	if err := ctx.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	validate := validation.Validator()
	if err := validate.Struct(req); err != nil {
		errorsMap := validation.CustomErrorMessages(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  "error",
			"errors":  errorsMap,
		})
	}
	task, err := c.TaskService.UpdateTask(ctx, req, taskId)
	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok && fiberErr.Code == fiber.StatusNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Task not found",
				"status":  "error",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update task",
			"status":  "error",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    response.ToTaskResponse(task),
		"message": "Task updated successfully",
		"status":  "success",
	})
}

// @Tags         Tasks
// @Summary      Delete a task
// @Description  Delete a task.
// @Produce      json
// @Param        taskId   path     string  true  "Task ID"
// @Success      200  {object}  response.TaskDeleteResponse
// @Failure      404  {object}  response.ErrorResponse  "Task not found"
// @Failure      500  {object}  response.ErrorResponse  "Internal server error"
// @Router       /tasks/{taskId} [delete]
func (c *TaskController) DeleteTask(ctx *fiber.Ctx) error {
	taskId := ctx.Params("taskId")
	err := c.TaskService.DeleteTask(ctx, taskId)
	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok && fiberErr.Code == fiber.StatusNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Task not found",
				"status":  "error",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete task",
			"status":  "error",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task deleted successfully",
		"status":  "success",
	})
}
