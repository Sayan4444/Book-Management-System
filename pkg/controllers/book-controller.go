package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Sayan4444/go-bookstore/pkg/database"
	"github.com/Sayan4444/go-bookstore/pkg/models"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var book models.Book
	if err := readJSON(r, &book); err != nil {
		writeJSONError(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	result := db.Create(&book)
	if result.Error != nil {
		writeJSONError(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var books []models.Book
	result := db.Find(&books)
	if result.Error != nil {
		writeJSONError(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, books)
}

func GetSingleBook(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]

	var book models.Book
	result := db.First(&book, id)

	if result.Error != nil {
		writeJSONError(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]

	var oldBook models.Book
	result := db.First(&oldBook, id)

	var updatedBook models.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		writeJSONError(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	if updatedBook.Name != "" {
		oldBook.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		oldBook.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		oldBook.Publication = updatedBook.Publication
	}

	if result.Error != nil {
		writeJSONError(w, "Failed to retrieve books", http.StatusNotFound)
		return
	}

	db.Save(&oldBook)

	writeJSON(w, http.StatusOK, oldBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]

	var book models.Book
	result := db.First(&book, id)

	if result.Error != nil {
		writeJSONError(w, "Invalid ID", http.StatusNotFound)
		return
	}

	result = db.Delete(&book, id)

	if result.Error != nil {
		writeJSONError(w, "Failed to delete books", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, struct{}{})
}
