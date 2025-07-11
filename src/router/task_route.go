package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(v1 fiber.Router, t service.TaskService) {
	taskController := controller.NewTaskController(t)

	task := v1.Group("/tasks")

	task.Get("/", taskController.GetTasks)

	task.Post("/", taskController.CreateTask)

	task.Put("/:taskId", taskController.UpdateTask)

	task.Delete("/:taskId", taskController.DeleteTask)

	task.Get("/:taskId", taskController.GetTaskByID)
}
