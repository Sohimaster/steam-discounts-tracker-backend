package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"steam_discounts_tracker-backend/db"
	"steam_discounts_tracker-backend/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	db.ConnectDB(mongoURI, dbName)

	router := mux.NewRouter()
	initializeRoutes(router)

	http.Handle("/", router)
	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func initializeRoutes(router *mux.Router) {
	router.HandleFunc("/subscriptions", handlers.AddSubscription).Methods("POST")
	router.HandleFunc("/subscriptions", handlers.ListSubscriptions).Methods("GET")
	router.HandleFunc("/subscriptions/{gameId}", handlers.RemoveSubscription).Methods("DELETE")
	router.HandleFunc("/search", handlers.SearchGames).Methods("GET")
	router.HandleFunc("/discounts", handlers.GetHotDiscounts).Methods("GET")
}
