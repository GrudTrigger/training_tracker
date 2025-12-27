package main

import (
	"context"
	"log"

	"github.com/GrudTrigger/training_tracker/backend/internal/app"
	"github.com/GrudTrigger/training_tracker/backend/internal/config"
)

const envPath = "./.env" // TODO: после заменить на /.env потому что при билде будет лежать в корне с .envß

func main() {
	err := config.Load(envPath)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
