package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Sayan4444/go-bookstore/pkg/database"
	"github.com/Sayan4444/go-bookstore/pkg/models"
	"github.com/Sayan4444/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	database.Connect()
	db := database.GetDB()
	db.AutoMigrate(&models.Book{})

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Printf("Server started")
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT not mentioned")
	}
	log.Fatal(http.ListenAndServe(":"+port, r))

}
