package main

import (
	"log"
	"net/http"
	"os"

	"github.com/2oofine/PaLuto/backend/internal/database"
	"github.com/2oofine/PaLuto/backend/internal/routes"
	"github.com/2oofine/PaLuto/backend/pkg/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	config.LoadEnv()
	database.Init()

	r := chi.NewRouter()
	r.Mount("/users", routes.UserRoutes(database.DB))

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // Default port if not set
	}

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
