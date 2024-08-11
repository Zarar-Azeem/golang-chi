package controllers

import (
	"fmt"
	"net/http"

	"github.com/Zarar-Azeem/golang-chi/pkg/models"
	"github.com/Zarar-Azeem/golang-chi/pkg/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books *models.Book
	result, err := books.GetAllBooks()

	if err != nil {
		fmt.Println("Error retrieving books:", err)
		return
	}
	jsonResponse := utils.ConverToJSON(w, result)
	w.Write(jsonResponse)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	// err := json.NewDecoder(r.Body).Decode(&book)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	name := r.FormValue("name")
	author := r.FormValue("author")
	publication := r.FormValue("publication")

	book = models.Book{
		Name:        name,
		Author:      author,
		Publication: publication,
	}

	book.AddBook(book)

	jsonResponse := utils.ConverToJSON(w, book)
	w.Write(jsonResponse)
}
