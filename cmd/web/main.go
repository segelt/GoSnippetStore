package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"snippetdemo/internal"
	"snippetdemo/internal/database/mongocl"
	"syscall"

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

	mongoParams := mongocl.DBParams{
		Username: os.Getenv("MONGODB_USER"),
		Password: os.Getenv("MONBODB_PWD"),
		Uri:      os.Getenv("MONGODB_URI"),
	}
	client, err := mongocl.NewMongoDB(mongoParams)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "mongodb.connect")
	}

	err = mongocl.SeedData(client)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "db seed")
	}

	errC := make(chan error, 1)
	go func() {
		jwtkey := os.Getenv("JWT_KEY")
		port := os.Getenv("PORT")
		app := &App{
			client:    client,
			secretKey: jwtkey,
			port:      port,
		}
		srv := app.NewServer()

		log.Println("Starting server")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errC <- err
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		log.Println("App shutdown.")

		defer func() {
			mongocl.GracefulShutdownDbConnection(client)
			stop()
			close(errC)
		}()
	}()

	return errC, nil
}
