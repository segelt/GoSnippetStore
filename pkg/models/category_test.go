package models

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var DBName string = "snippetdbtest"

func setupTest(t *testing.T) func(t *testing.T) {

	uri := "mongodb://localhost:27017"
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(uri)

	var merr error
	client, merr = mongo.Connect(ctx, clientOptions)

	if merr != nil {
		t.Fatal(merr)
	}

	if terr := client.Ping(ctx, readpref.Primary()); terr != nil {
		t.Error("mongodb ping failed..")
	}

	return func(t *testing.T) {
		t.Log("teardown test")
	}
}

func TestCategoriesFilterId(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	repo := &CategoryModel{Client: client, DBName: DBName}
	var categoryid int = 1
	filter := CategoryFilter{
		CategoryId: &categoryid,
	}

	results, err := repo.Filter(filter)

	if err != nil {
		t.Fatalf(err.Error())
	}

	expectedcount := 1
	if len(*results) != expectedcount {
		t.Fatalf("Expected category count does not match")
	}

	expectedDesc := "testcategory1"
	if (*results)[0].Description != expectedDesc {
		t.Fatalf("Expected category description does not match")
	}
}

func TestCategoriesFilterDescSort(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	repo := &CategoryModel{Client: client, DBName: DBName}
	var sortby string = "description"
	var sortdirection string = "asc"
	var description string = "testc"
	filter := CategoryFilter{
		SortBy:        &sortby,
		SortDirection: &sortdirection,
		Description:   &description,
	}

	results, err := repo.Filter(filter)

	if err != nil {
		t.Fatalf(err.Error())
	}

	expectedcount := 2
	if len(*results) != expectedcount {
		t.Fatalf("Expected category count does not match")
	}

	expectedDescFirst := "testcategory1"
	if (*results)[0].Description != expectedDescFirst {
		t.Fatalf("Expected category description does not match")
	}

	expectedDescSecond := "testcategory2"
	if (*results)[1].Description != expectedDescSecond {
		t.Fatalf("Expected category description does not match")
	}
}

func TestCategoriesFilterNoMatch(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	repo := &CategoryModel{Client: client, DBName: DBName}
	var sortby string = "description"
	var sortdirection string = "asc"
	var description string = "testccc"
	filter := CategoryFilter{
		SortBy:        &sortby,
		SortDirection: &sortdirection,
		Description:   &description,
	}

	results, err := repo.Filter(filter)

	if err != nil {
		t.Fatalf(err.Error())
	}

	expectedcount := 0
	if len(*results) != expectedcount {
		t.Fatalf("Expected category count does not match")
	}
}

func TestCategoriesSingle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	repo := &CategoryModel{Client: client, DBName: DBName}

	results, err := repo.Single(1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	expectedDescription := "testcategory1"
	if results.Description != expectedDescription {
		t.Fatalf("Expected category description does not match")
	}
}

func TestCategoriesSingleNonMatch(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	repo := &CategoryModel{Client: client, DBName: DBName}

	results, err := repo.Single(1111111)

	if err == nil {
		t.Fatalf(err.Error())
	}

	if results != nil {
		t.Fatalf("no category should have been returned")
	}
}

func TestCategoriesByUserId(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	repo := &CategoryModel{Client: client, DBName: DBName}

	results, err := repo.ByUser("632655b353adec83f7f2d6a5")

	if err != nil {
		t.Fatalf(err.Error())
	}

	expectedcount := 2
	if len(results) != expectedcount {
		t.Fatalf("Expected category count does not match")
	}
}
