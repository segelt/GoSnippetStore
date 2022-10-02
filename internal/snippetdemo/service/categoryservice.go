package service

import (
	"context"
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

func (svc *CategoryService) AddCategory(ctx context.Context, categoryId int, description string) error {
	err := svc.categories.Upsert(ctx, categoryId, description)

	return err
}

func (svc *CategoryService) GetCategoryById(ctx context.Context, categoryId int) (*models.Category, error) {

	targetCategory, err := svc.categories.Single(ctx, categoryId)

	if err != nil {
		return nil, err
	}

	return targetCategory, nil
}

func (svc *CategoryService) GetCategories(ctx context.Context, filter models.CategoryFilter) (*[]models.Category, error) {
	results, err := svc.categories.Filter(ctx, filter)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func NewCategoryService(client *mongo.Client, DBName string) *CategoryService {
	return &CategoryService{categories: models.CategoryModel{Client: client, DBName: DBName}}
}
