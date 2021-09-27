package main

import (
	"log"

	"github.com/brandonmillerio/nuboverflow_users/repository/db"
	"github.com/brandonmillerio/nuboverflow_users/repository/user"
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
	_ = user.New(userStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
