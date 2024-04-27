package handlers

import (
	"encoding/json"
	"net/http"
)

func GetHotDiscounts(w http.ResponseWriter, r *http.Request) {
	// This should fetch discount information, possibly from the Steam API
	// Placeholder response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("List of current hottest discounts")
}
