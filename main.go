package main

import (
	credential "go-api/module"
	routes_actors "go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login", credential.Send_token)
	v1 := router.Group("/api/v1")
	{
		routes_actors.RegisterActorRoutes(v1) // Appeler les routes des acteurs
	}
	router.Run(":8080")
}
