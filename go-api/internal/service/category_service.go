package service

import (
	"github.com/gabrsobral/imersao17/goapi/internal/database"
	"github.com/gabrsobral/imersao17/goapi/internal/entity"
)

type CategoryService struct {
	CategoryDb database.CategoryDb
}

func NewCategoryService(categoryDb *database.CategoryDb) *CategoryService {
	return &CategoryService{CategoryDb: *categoryDb}
}

func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := cs.CategoryDb.GetCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cs *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)
	_, err := cs.CategoryDb.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cs *CategoryService) GetCategoryById(categoryId string) (*entity.Category, error) {
	category, err := cs.CategoryDb.GetCategoryById(categoryId)

	if err != nil {
		return nil, err
	}

	return category, err
}
