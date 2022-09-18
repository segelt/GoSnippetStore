package service

import (
	"context"
	"fmt"
	"log"
	"snippetdemo/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Category models.Category

type CategoryService struct {
	Client *mongo.Client
}

type CategoryFilter struct {
	CategoryId    *int
	Description   *string
	SortBy        *string
	SortDirection *string
	Count         *int
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

func (svc *CategoryService) GetCategories(filter CategoryFilter) (*[]Category, error) {
	coll := svc.Client.Database("snippetdb").Collection("categories")

	qry := bson.D{}
	if filter.CategoryId != nil {
		f := bson.E{Key: "categoryId", Value: *filter.CategoryId}
		qry = append(qry, f)
	}

	if filter.Description != nil {
		var descfilter string
		descfilter = fmt.Sprintf("/.*%s.*/", *filter.Description)
		f := bson.E{Key: "description", Value: bson.E{
			Key: "$regex",
			Value: primitive.Regex{
				Pattern: ".*" + descfilter + ".*",
				Options: "i",
			},
		}}
		qry = append(qry, f)
	}

	var sortOptions *options.FindOptions
	var sortFilter bson.D
	if filter.SortBy != nil {
		sortDir := -1
		if *filter.SortDirection == "asc" {
			sortDir = 1
		}

		sortFilter = make(bson.D, 0)
		switch *filter.SortBy {
		case "description":
			sortFilter = append(sortFilter, bson.E{"description", sortDir})
		case "title":
			sortFilter = append(sortFilter, bson.E{"title", sortDir})
		case "count":
			panic("Not implemented..")
		}

		sortOptions = options.Find().SetSort(sortFilter)
	}

	cursor, err := coll.Find(context.TODO(), filter, sortOptions)

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return nil, nil
}

func NewCategoryService(client *mongo.Client) *CategoryService {
	return &CategoryService{Client: client}
}
