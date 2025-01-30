package handlers

import (
	"encoding/json"
	"net/http"
)

func GetArtPrompt(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"art_prompt": "Circle"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ArtPromptHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetArtPrompt(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
