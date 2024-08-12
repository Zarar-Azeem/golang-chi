package routes

import (
	"github.com/Zarar-Azeem/golang-chi/pkg/controllers"
	"github.com/go-chi/chi/v5"
)

func RegisterBookStoreRoutes(router chi.Router) {
	router.Get("/", controllers.GetBooks)
	router.Get("/{id}", controllers.GetBook)
	router.Post("/", controllers.CreateBook)
	router.Put("/{id}", controllers.UpdateBook)
	router.Delete("/{id}", controllers.DeleteBook)
}
