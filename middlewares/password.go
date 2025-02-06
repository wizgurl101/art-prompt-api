package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	db "art-prompt-api/db"
	models "art-prompt-api/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func HashPasswordMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		hashedPassword, err := HashPassword(user.Password)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		user.Password = hashedPassword

		ctx := context.WithValue(r.Context(), models.UserContextKey, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func VerifyPasswordMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		found_user, err := db.GetUser(user.Email)
		if err != nil {
			http.Error(w, "Email not found", http.StatusInternalServerError)
			return
		}

		err = VerifyPassword(found_user.Password, user.Password)
		if err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), models.UserContextKey, found_user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
