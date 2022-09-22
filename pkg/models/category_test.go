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

func TestCategoriesByUserId(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	repo := &CategoryModel{Client: client}

	_, err := repo.ByUser("632655b353adec83f7f2d6a5")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
