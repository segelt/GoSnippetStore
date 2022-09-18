package service

import (
	"context"
	"log"
	"snippetdemo/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Category models.Category

type CategoryService struct {
	Client *mongo.Client
}

func (svc *CategoryService) AddCategory(categoryId int, description string) error {
	coll := svc.Client.Database("snippetdb").Collection("categories")

	category := bson.D{{"categoryId", categoryId}, {"description", description}}

	_, err := coll.InsertOne(context.TODO(), category)
	return err
}

func (svc *CategoryService) GetCategoryById(categoryId int) (*Category, error) {
	coll := svc.Client.Database("snippetdb").Collection("categories")

	var targetCategory *Category
	err = coll.FindOne(context.TODO(), bson.M{"categoryId": categoryId}).Decode(targetCategory)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return targetCategory, nil
}
