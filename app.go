package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	// Creates a gloval instance called 'logger'
	logger := logrus.New()

	// Formats logs into parsable json
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}

	// Create a new Mux object that implements the Router interface
	r := chi.NewRouter()

	// Base middleware stack
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/fibo", Fibo)
	})

	// Start server
	fmt.Println("Start listening http at port 8090...")
	err := http.ListenAndServe(":8090", r)
	if err != nil {
		log.Print(err.Error())
		return
	}
}
