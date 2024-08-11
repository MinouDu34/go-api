package routes_actors

import (
	"github.com/gin-gonic/gin"

	actor_crud "go-api/crud"
)

// RegisterActorRoutes définit les routes pour les acteurs
func RegisterActorRoutes(r *gin.RouterGroup) {
	r.GET("/actors/:id", actor_crud.GetActorById)
	r.GET("/actors", actor_crud.GetActors)
	r.POST("/ ", actor_crud.CreateActor)
	r.DELETE("/actors/:id", actor_crud.DelActorById)
}
