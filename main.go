package main

import (
	"art-prompt-api/routes"
	"fmt"
	"net/http"
)

func main() {
	routes := routes.RegisterRoutes()
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", routes); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
