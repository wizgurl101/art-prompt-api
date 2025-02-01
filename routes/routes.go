package routes

import (
	"art-prompt-api/handlers"
	middlewares "art-prompt-api/middlewares"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.Handle("/prompt", middlewares.AuthMiddleware(http.HandlerFunc(handlers.ArtPromptHandler)))
	mux.HandleFunc("/user", handlers.UserHandler)
	return mux
}
