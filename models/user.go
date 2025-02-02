package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type contextKey string

const UserContextKey contextKey = "user"
