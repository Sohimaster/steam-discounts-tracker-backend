package handlers

import (
	"encoding/json"
	"net/http"
)

// SearchGames handles GET requests to search for games by name
func SearchGames(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("name")
	// This should call the Steam API and return results
	// Placeholder response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Search results for: " + query)
}
