package actor_crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActors(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Liste des acteurs en mode route !",
	})
}

func CreateActor(c *gin.Context) {
	var newActor struct {
		Name   string `json:"name" binding:"required"`
		Age    int    `json:"age" binding:"required"`
		Gender string `json:"gender" binding:"required"`
	}

	// Lier les données JSON de la requête au struct
	if err := c.ShouldBindJSON(&newActor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Répondre avec les données reçues
	c.JSON(http.StatusOK, gin.H{
		"message": "Acteur créé avec succès",
		"actor":   newActor,
	})
}

func GetActorById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "GetActorById " + id + " Called"})
}

func DelActorById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "DelActorById " + id + " Called"})
}
