package mongocl

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Repo struct {
	Client *mongo.Client
}

func (s *Repo) setupConnection(uri string) error {
	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Print(evt.Command)
		},
	}

	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(uri).SetMonitor(cmdMonitor)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	s.Client = client
	return nil
}

func (s *Repo) Initialize(uri string) error {
	err := s.setupConnection(uri)

	if err != nil {
		log.Fatal(err)
		log.Fatal("Could not connect to db. Terminating app")
		return err
	}

	err = s.seedCategories()

	if err != nil {
		log.Fatal(err)
		log.Fatal("Could not seed categories.")
		return err
	}

	log.Println("Data seeded.")

	return nil
}

func (s *Repo) seedCategories() error {
	coll := s.Client.Database("snippetdb").Collection("categories")

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

func (s Repo) GracefulShutdownDbConnection() error {
	fmt.Println("Disconnecting")
	err := s.Client.Disconnect(context.TODO())

	return err
}
