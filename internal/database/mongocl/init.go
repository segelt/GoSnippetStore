package mongocl

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Repo struct {
	Client *mongo.Client
}

func (s *Repo) setupConnection(uri string) error {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(uri)
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

	return nil
}

func (s Repo) GracefulShutdownDbConnection() error {
	fmt.Println("Disconnecting")
	err := s.Client.Disconnect(context.TODO())

	return err
}
