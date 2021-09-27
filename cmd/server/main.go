package main

import (
	"log"

	"github.com/brandomillerio/nuboverflow_users/internal/db"
	"github.com/brandomillerio/nuboverflow_users/internal/rocket"
)

func Run() error {
	//Responsible for initializing and starting the gRPC server.
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}
	_ = rocket.New(rocketStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
