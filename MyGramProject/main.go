// main.go
package main

import (
	"MyGramProject/handlers"
	"MyGramProject/repositories"
	"MyGramProject/routers"
	"MyGramProject/services"
	"fmt"
	"log"
)

func main() {
	// Initialize database connection
	db, err := repositories.InitDatabase()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	router := routers.SetupRouter(userHandler)

	// Start the server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	router.Run(port)
}
