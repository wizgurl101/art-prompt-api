package handlers

import (
	"art-prompt-api/db"
	models "art-prompt-api/models"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	new_user, isOk := r.Context().Value(models.UserContextKey).(models.User)
	if !isOk {
		http.Error(w, "Invalid request payload", http.StatusInternalServerError)
		return
	}

	collection := db.GetCollection("users")
	_, err := collection.InsertOne(r.Context(), new_user)
	if err != nil {
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Account created"})
}

// func GetUser(w http.ResponseWriter, r *http.Request) models.User {
// 	var user models.User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return models.User{}
// 	}

// 	collection := db.GetCollection("users")
// 	foundUser, err := collection.FindOne(r.Context(), user)
// 	if err != nil {
// 		http.Error(w, "Failed to find user", http.StatusInternalServerError)
// 		return models.User{}
// 	}

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(map[string]string{"message": msg})
// }

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateUser(w, r)
	default:
		http.Error(w, "404 Not Found", http.StatusMethodNotAllowed)
	}
}
