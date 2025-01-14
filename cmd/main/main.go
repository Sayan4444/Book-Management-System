package main

import (
	"log"
	"net/http"

	"github.com/Sayan4444/go-bookstore/pkg/database"
	"github.com/Sayan4444/go-bookstore/pkg/models"
	"github.com/Sayan4444/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	db := database.GetDB()
	db.AutoMigrate(&models.Book{})

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Printf("Server started")
	log.Fatal(http.ListenAndServe("localhost:4000", r))

}
