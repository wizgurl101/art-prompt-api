package main

import (
	"art-prompt-api/db"
	"art-prompt-api/routes"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	db.InitializeRedis()

	routes := routes.RegisterRoutes()
	fmt.Println("Art Prompt API Server is running")
	if err := http.ListenAndServe(":8080", routes); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
