package handlers

import (
	"art-prompt-api/db"
	"net/http"
)

func clearAllKeys(w http.ResponseWriter, r *http.Request) {
	err := db.ClearAllKeys()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RedisHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		clearAllKeys(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
