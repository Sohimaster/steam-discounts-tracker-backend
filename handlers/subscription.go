package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// AddSubscription handles POST requests to add a new game to the subscription list
func AddSubscription(w http.ResponseWriter, r *http.Request) {
	// This would typically involve decoding a JSON request body and updating a database
	// Placeholder response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Game added to subscription list")
}

func ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Subscription List")
}

// RemoveSubscription handles DELETE requests to remove a game from the subscription list
func RemoveSubscription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	// This would typically involve removing an item from a database
	// Placeholder response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Game removed from subscription list: " + gameId)
}
