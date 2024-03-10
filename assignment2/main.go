package main

import (
	"assignment2/repositories"
	"assignment2/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repositories.InitDatabase()

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	routers.OrderRoutes(router, db)

	router.Run()
}
