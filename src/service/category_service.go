package service

import (
	"app/src/model"
	"app/src/utils"
	"app/src/validation"

	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryService interface {
	GetCategories(c *fiber.Ctx, params *validation.QueryCategory) ([]model.Category, int64, error)
	GetCategoryByID(c *fiber.Ctx, id string) (*model.Category, error)
	CreateCategory(c *fiber.Ctx, req *validation.CreateCategory) (*model.Category, error)
	UpdateCategory(c *fiber.Ctx, req *validation.UpdateCategory, id string) (*model.Category, error)
	DeleteCategory(c *fiber.Ctx, id string) error
}

type categoryService struct {
	Log      *logrus.Logger
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewCategoryService(db *gorm.DB, validate *validator.Validate) CategoryService {
	return &categoryService{
		Log:      utils.Log,
		DB:       db,
		Validate: validate,
	}
}

func (s *categoryService) GetCategories(c *fiber.Ctx, params *validation.QueryCategory) ([]model.Category, int64, error) {
	var categories []model.Category
	var total int64

	query := s.DB.Model(&model.Category{})

	if params.Search != "" {
		query = query.Where("name ILIKE ?", "%"+params.Search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		s.Log.Errorf("Failed to count categories: %v", err)
		return nil, 0, err
	}

	query = query.Order("created_at DESC")
	query = query.Offset(params.Page * params.Limit).Limit(params.Limit)

	if err := query.Find(&categories).Error; err != nil {
		s.Log.Errorf("Failed to get categories: %v", err)
		return nil, 0, err
	}

	return categories, total, nil
}

func (s *categoryService) GetCategoryByID(c *fiber.Ctx, id string) (*model.Category, error) {
	category := new(model.Category)

	result := s.DB.WithContext(c.Context()).First(category, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fiber.NewError(fiber.StatusNotFound, "Category not found")
	}

	return category, nil
}

func (s *categoryService) CreateCategory(c *fiber.Ctx, req *validation.CreateCategory) (*model.Category, error) {
	var category model.Category

	category.Name = req.Name

	if err := s.DB.Create(&category).Error; err != nil {
		s.Log.Errorf("Failed to create category: %v", err)
		return nil, err
	}

	return &category, nil
}

func (s *categoryService) UpdateCategory(c *fiber.Ctx, req *validation.UpdateCategory, id string) (*model.Category, error) {
	// find category by id, when not found, return error
	category, err := s.GetCategoryByID(c, id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Category not found")
	}

	if err := s.DB.Model(category).Where("id = ?", id).Updates(req).Error; err != nil {
		s.Log.Errorf("Failed to update category: %v", err)
		return nil, err
	}

	return category, nil
}

func (s *categoryService) DeleteCategory(c *fiber.Ctx, id string) error {
	// find category by id, when not found, return error
	_, err := s.GetCategoryByID(c, id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Category not found")
	}

	if err := s.DB.Delete(&model.Category{}, "id = ?", id).Error; err != nil {
		s.Log.Errorf("Failed to delete category: %v", err)
		return err
	}

	return nil
}
