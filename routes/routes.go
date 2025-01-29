package routes

import (
	"art-prompt-api/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/prompt", handlers.ArtPromptHandler)
}
