package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(v1 fiber.Router, c service.CategoryService) {
	categoryController := controller.NewCategoryController(c)

	category := v1.Group("/categories")

	category.Get("/", categoryController.GetCategories)
	category.Post("/", categoryController.CreateCategory)
	category.Put("/:categoryId", categoryController.UpdateCategory)
	category.Delete("/:categoryId", categoryController.DeleteCategory)
	category.Get("/:categoryId", categoryController.GetCategoryByID)
}
