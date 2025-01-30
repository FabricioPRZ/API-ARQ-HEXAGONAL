package routes

import (
	"API-HEXAGONAL/src/shinobi/dependencies"
	"github.com/gin-gonic/gin"
)

// Configura las rutas para los shinobis
func SetupShinobiRoutes(router *gin.Engine, deps *dependencies.ShinobiDependencies) {
	// Grupo de rutas para los shinobis
	shinobiGroup := router.Group("/v1/shinobis")
	{
		shinobiGroup.POST("", deps.CreateShinobiController.Run)
		shinobiGroup.PUT("/:id", deps.UpdateShinobiController.Run)
		shinobiGroup.GET("", deps.GetAllShinobisController)
		shinobiGroup.DELETE("/:id", deps.DeleteShinobiController.Run)
	}
}
