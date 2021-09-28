package main

import (
	"log"

	"github.com/brandonmillerio/nuboverflow_users/internal/db"
	"github.com/brandonmillerio/nuboverflow_users/internal/transport/grpc"
	"github.com/brandonmillerio/nuboverflow_users/internal/user"
)

func Run() error {
	//Responsible for initializing and starting the gRPC server.
	userStore, err := db.New()
	if err != nil {
		return err
	}
	err = userStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}
	
	usrService := user.New(userStore)
	usrHandler := grpc.New(usrService)


    if err := usrHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
