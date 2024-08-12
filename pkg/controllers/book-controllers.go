package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
	message := "Retrieved Successfully"

	jsonResponse := utils.ConverToJSON(w, result, message)
	w.Write(jsonResponse)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	book, err := models.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	message := "Retrieved Successfully"

	jsonResponse := utils.ConverToJSON(w, book, message)

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
	message := "Created Successfully"
	jsonResponse := utils.ConverToJSON(w, book, message)
	w.Write(jsonResponse)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	success := false
	id, _ := strconv.Atoi(r.PathValue("id"))
	_, err := models.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	models.DeleteByID(id)
	success = true
	message := "Deleted Successfully"
	jsonRepsonse := utils.ConverToJSON(w, success, message)
	w.Write(jsonRepsonse)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	book , err := models.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	name := r.FormValue("name")
	author := r.FormValue("author")
	publication := r.FormValue("publication")

	book.Name = name
	book.Author = author
	book.Publication = publication

	models.UpdateBook(book)
	message := "Updated Successfully"
	jsonRepsonse := utils.ConverToJSON(w,  book ,message)
	w.Write(jsonRepsonse)
}
