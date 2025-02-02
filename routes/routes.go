package routes

import (
	"art-prompt-api/handlers"
	middlewares "art-prompt-api/middlewares"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/login", middlewares.VerifyPasswordMiddleware(http.HandlerFunc(handlers.LoginHandler)))
	mux.Handle("/prompt", middlewares.AuthMiddleware(http.HandlerFunc(handlers.ArtPromptHandler)))
	mux.Handle("/create-user", middlewares.HashPasswordMiddleware(http.HandlerFunc(handlers.UserHandler)))
	return mux
}
