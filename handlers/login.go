package handlers

import (
	auth "art-prompt-api/middlewares"
	models "art-prompt-api/models"
	"encoding/json"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	user, isOk := r.Context().Value(models.UserContextKey).(models.User)
	if !isOk {
		http.Error(w, "Invalid request payload", http.StatusInternalServerError)
		return
	}

	token, err := auth.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		LoginUser(w, r)
	default:
		http.Error(w, "404 Not Found", http.StatusMethodNotAllowed)
	}
}
