package main

import (
	"log"
	"net/http"

	genexercise "github.com/GrudTrigger/training_tracker/backend/gen/exercise"
	genhttp "github.com/GrudTrigger/training_tracker/backend/gen/http/exercise/server"
	"github.com/GrudTrigger/training_tracker/backend/internal/service/exersice"
	goahttp "goa.design/goa/v3/http"
)

func main() {
	exerciseSvc := exersice.NewExersiceService()
	endpoints := genexercise.NewEndpoints(exerciseSvc)

	mux := goahttp.NewMuxer()

	handler := genhttp.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)

	genhttp.Mount(mux, handler)

	port := "8080"
	server := &http.Server{Addr: ":" + port, Handler: mux}

	for _, mount := range handler.Mounts {
		log.Printf("%q mounted on %s %s", mount.Method, mount.Verb, mount.Pattern)
	}

	log.Printf("Starting concerts service on :%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
