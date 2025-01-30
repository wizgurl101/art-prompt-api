package main

import (
	"art-prompt-api/routes"
	"fmt"
	"net/http"
)

func main() {
	routes := routes.RegisterRoutes()
	fmt.Println("Art Prompt API Server is running on port 5000")
	if err := http.ListenAndServe(":5000", routes); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
