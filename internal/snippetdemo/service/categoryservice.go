package service

import (
	"snippetdemo/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryService struct {
	categories models.CategoryModel
}

type CategoryFilter struct {
	CategoryId    *int
	Description   *string
	SortBy        *string
	SortDirection *string
	Count         *int
}

func (svc *CategoryService) AddCategory(categoryId int, description string) error {
	err := (&svc.categories).Upsert(categoryId, description)

	return err
}

func (svc *CategoryService) GetCategoryById(categoryId int) (*models.Category, error) {

	targetCategory, err := (&svc.categories).Single(categoryId)

	if err != nil {
		return nil, err
	}

	return targetCategory, nil
}

func (svc *CategoryService) GetCategories(filter models.CategoryFilter) (*[]models.Category, error) {
	results, err := (&svc.categories).Filter(filter)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func NewCategoryService(client *mongo.Client) *CategoryService {
	return &CategoryService{categories: models.CategoryModel{Client: client}}
}
