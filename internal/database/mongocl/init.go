package mongocl

import (
	"context"
	"fmt"
	"log"
	"snippetdemo/internal"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// type Repo struct {
// 	Client *mongo.Client
// }

func NewMongoDB(uri string) (*mongo.Client, error) {
	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Print(evt.Command)
		},
	}

	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(uri).SetMonitor(cmdMonitor)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "mongodb.connect")
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "mongodb.ping")
	}

	// err = s.seedCategories()

	// if err != nil {
	// 	log.Fatal(err)
	// 	log.Fatal("Could not seed categories.")
	// 	return err
	// }

	// log.Println("Data seeded.")

	return client, nil
}

func SeedData(client *mongo.Client) error {
	dbname := "snippetdb"
	coll := client.Database(dbname).Collection("categories")

	filter_testcategory1 := bson.D{{"categoryId", 1}}
	update_testcategory1 := bson.D{{"$set", bson.D{{"categoryId", 1}, {"description", "testcategory1"}}}}

	filter_testcategory2 := bson.D{{"categoryId", 2}}
	update_testcategory2 := bson.D{{"$set", bson.D{{"categoryId", 2}, {"description", "testcategory2"}}}}

	opts := options.Update().SetUpsert(true)

	_, err := coll.UpdateOne(context.TODO(), filter_testcategory1, update_testcategory1, opts)
	if err != nil {
		return err
	}

	_, err = coll.UpdateOne(context.TODO(), filter_testcategory2, update_testcategory2, opts)
	if err != nil {
		return err
	}

	return nil
}

func GracefulShutdownDbConnection(client *mongo.Client) error {
	fmt.Println("Disconnecting")
	err := client.Disconnect(context.TODO())

	return err
}
