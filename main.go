package main

import (
	routes_actors "go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		routes_actors.RegisterActorRoutes(v1) // Appeler les routes des acteurs
	}
	router.Run(":8080")
}
