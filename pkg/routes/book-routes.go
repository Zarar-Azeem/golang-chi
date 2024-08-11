package routes

import (
	"github.com/Zarar-Azeem/golang-chi/pkg/controllers"
	"github.com/go-chi/chi/v5"
)

func RegisterBookStoreRoutes(router chi.Router) {
	router.Get("/", controllers.GetBooks)
	router.Post("/", controllers.CreateBook)
}
