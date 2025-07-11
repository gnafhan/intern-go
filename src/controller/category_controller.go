package controller

import (
	"app/src/service"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

// @Tags         Categories
// @Summary      Get all categories
// @Description  Retrieve all categories.
// @Produce      json
// @Param        page     query     int     false   "Page number"  default(0)
// @Param        limit    query     int     false   "Maximum number of categories"    default(10)
// @Param        search   query     string  false  "Search by name"
// @Success      200  {object}  example.GetCategoriesResponse
// @Failure      500  {object}  example.InternalServerError  "Internal server error"
// @Router       /categories [get]
func (c *CategoryController) GetCategories(ctx *fiber.Ctx) error {
	query := &validation.QueryCategory{
		Page:   ctx.QueryInt("page", 1),
		Limit:  ctx.QueryInt("limit", 10),
		Search: ctx.Query("search", ""),
	}

	categories, total, err := c.CategoryService.GetCategories(ctx, query)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get categories",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    categories,
		"total":   total,
		"message": "Categories retrieved successfully",
		"status":  "success",
	})
}

// @Tags         Categories
// @Summary      Get a category by ID
// @Description  Retrieve a category by ID.
// @Produce      json
// @Param        categoryId   path     string  true  "Category ID"
// @Success      200  {object}  example.GetCategoryByIDResponse
// @Failure      404  {object}  example.NotFound  "Category not found"
// @Router       /categories/{categoryId} [get]
func (c *CategoryController) GetCategoryByID(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")

	category, err := c.CategoryService.GetCategoryByID(ctx, categoryId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    category,
		"message": "Category retrieved successfully",
		"status":  "success",
	})
}

// @Tags         Categories
// @Summary      Create a category
// @Description  Create a category.
// @Produce      json
// @Param        category   body     validation.CreateCategory  true  "Category"
// @Success      201  {object}  example.CreateCategoryResponse
// @Failure      400  {object}  example.BadRequest  "Invalid request body"
// @Failure      500  {object}  example.InternalServerError  "Internal server error"
// @Router       /categories [post]
func (c *CategoryController) CreateCategory(ctx *fiber.Ctx) error {
	req := new(validation.CreateCategory)

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  "error",
		})
	}

	validate := validation.Validator()

	if err := validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  "error",
		})
	}

	category, err := c.CategoryService.CreateCategory(ctx, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create category",
			"status":  "error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    category,
		"message": "Category created successfully",
		"status":  "success",
	})
}

// @Tags         Categories
// @Summary      Update a category
// @Description  Update a category.
// @Produce      json
// @Param        categoryId   path     string  true  "Category ID"
// @Param        category   body     validation.UpdateCategory  true  "Category"
// @Success      200  {object}  example.UpdateCategoryResponse
// @Failure      400  {object}  example.BadRequest  "Invalid request body"
// @Failure      500  {object}  example.InternalServerError  "Internal server error"
// @Router       /categories/{categoryId} [put]
func (c *CategoryController) UpdateCategory(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")

	req := new(validation.UpdateCategory)

	if err := ctx.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	validate := validation.Validator()

	if err := validate.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	category, err := c.CategoryService.UpdateCategory(ctx, req, categoryId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    category,
		"message": "Category updated successfully",
		"status":  "success",
	})
}

// @Tags         Categories
// @Summary      Delete a category
// @Description  Delete a category.
// @Produce      json
// @Param        categoryId   path     string  true  "Category ID"
// @Success      200  {object}  example.DeleteCategoryResponse
// @Failure      404  {object}  example.NotFound  "Category not found"
// @Failure      500  {object}  example.InternalServerError  "Internal server error"
// @Router       /categories/{categoryId} [delete]
func (c *CategoryController) DeleteCategory(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")

	if err := c.CategoryService.DeleteCategory(ctx, categoryId); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category deleted successfully",
		"status":  "success",
	})
}
