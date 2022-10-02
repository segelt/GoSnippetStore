package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"snippetdemo/internal"
	"snippetdemo/internal/database/mongocl"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	log.SetOutput(os.Stdout)

	errC, err := run()
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func run() (<-chan error, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "env.Load")
	}

	mongoUri := os.Getenv("MONGODB_URI")
	client, err := mongocl.NewMongoDB(mongoUri)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "mongodb.connect")
	}

	err = mongocl.SeedData(client)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "db seed")
	}

	jwtkey := os.Getenv("JWT_KEY")
	srv := &Server{client: client, secretKey: jwtkey}
	err = srv.StartServer()
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "server.start")
	}

	fmt.Println("Started server")

	errC := make(chan error, 1)
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		fmt.Println("App shutdown.")
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			mongocl.GracefulShutdownDbConnection(client)
			stop()
			cancel()
			close(errC)
		}()
	}()

	return errC, nil
}
