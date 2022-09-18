package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID            primitive.ObjectID `json:"innerId" bson:"_id,omitempty"`
	CategoryID    int                `json:"categoryId" bson:"categoryId"`
	Description   int                `json:"description" bson:"description"`
	TotalSnippets int
}

type CategoryService interface {
	AddCategory(categoryId int, description string) error
	GetCategoryById(categoryId int) (Category, error)
}
