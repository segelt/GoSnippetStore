package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Category struct {
	ID            primitive.ObjectID `json:"innerId" bson:"_id,omitempty"`
	CategoryID    int                `json:"categoryId" bson:"categoryId"`
	Description   string             `json:"description" bson:"description"`
	TotalSnippets int
}

type CategoryService interface {
	AddCategory(categoryId int, description string) error
	GetCategoryById(categoryId int) (Category, error)
}

type CategoryModel struct {
	Client *mongo.Client
}

type CategoryFilter struct {
	CategoryId    *int
	Description   *string
	SortBy        *string
	SortDirection *string
	Count         *int
}

type NestedCategoryInfo struct {
	CategoryId          int    `bson:"categoryId"`
	CategoryDescription string `bson:"categoryDescription"`
}

type ByUserResult struct {
	CategoryInfo NestedCategoryInfo `bson:"_id"`
	// categoryInfo interface{} `bson:"_id"`
	Count int `bson:"amount"`
}

func (c *CategoryModel) Filter(filter CategoryFilter) (*[]Category, error) {
	coll := c.Client.Database("snippetdb").Collection("categories")

	qry := bson.D{}
	if filter.CategoryId != nil {
		f := bson.E{Key: "categoryId", Value: *filter.CategoryId}
		qry = append(qry, f)
	}

	if filter.Description != nil {
		f := bson.E{Key: "description",
			Value: bson.D{{"$regex",
				primitive.Regex{
					Pattern: *filter.Description,
					Options: "i"}},
			},
		}

		qry = append(qry, f)
	}

	var sortOptions *options.FindOptions
	var sortFilter bson.D
	if filter.SortBy != nil {
		sortDir := -1
		if filter.SortDirection != nil && *filter.SortDirection == "asc" {
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

	cursor, _ := coll.Find(context.TODO(), qry, sortOptions)

	var results []Category
	if err := cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &results, nil
}

func (c *CategoryModel) Single(categoryId int) (*Category, error) {
	coll := c.Client.Database("snippetdb").Collection("categories")

	var targetCategory *Category
	err := coll.FindOne(context.TODO(), bson.M{"categoryId": categoryId}).Decode(targetCategory)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return targetCategory, nil
}

func (c *CategoryModel) ByUser(userid string) ([]Category, error) {

	snippetsCol := c.Client.Database("snippetdb").Collection("snippets")
	cursor, err := snippetsCol.Aggregate(context.TODO(), bson.A{
		bson.D{{"$match", bson.D{{"UserId", userid}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "categories"},
					{"localField", "category"},
					{"foreignField", "categoryId"},
					{"as", "category"},
				},
			},
		},
		bson.D{{"$unwind", bson.D{{"path", "$category"}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id",
						bson.D{
							{"categoryId", "$category.categoryId"},
							{"categoryDescription", "$category.description"},
						},
					},
					{"amount", bson.D{{"$sum", 1}}},
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	var groupingResults []ByUserResult
	err = cursor.All(context.TODO(), &groupingResults)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return nil, err
}

func (c *CategoryModel) Upsert(categoryId int, description string) error {
	coll := c.Client.Database("snippetdb").Collection("categories")

	filter_testcategory := bson.D{{"categoryId", categoryId}}
	update_testcategory := bson.D{{"$set", bson.D{{"categoryId", categoryId}, {"description", description}}}}

	opts := options.Update().SetUpsert(true)

	_, err := coll.UpdateOne(context.TODO(), filter_testcategory, update_testcategory, opts)
	if err != nil {
		return err
	}

	return nil
}
