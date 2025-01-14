package routes

import (
	"github.com/Sayan4444/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetAllBooks).Methods("GET")

	router.HandleFunc("/book/{id}", controllers.GetSingleBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
