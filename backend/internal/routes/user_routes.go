package routes

import (
	"net/http"

	"github.com/2oofine/PaLuto/backend/internal/controllers"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	uc := &controllers.UserController{
		DB: db,
	}

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) { return })
	r.Post("/register", uc.RegisterUser)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { return })
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) { return })

	return r
}
