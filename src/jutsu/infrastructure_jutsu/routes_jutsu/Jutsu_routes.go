package routes_jutsu

import (
	"API-HEXAGONAL/src/jutsu/dependencies_jutsu"
	"github.com/gin-gonic/gin"
)

// Configura las rutas para los jutsus
func SetupJutsuRoutes(router *gin.Engine, deps *dependencies_jutsu.JutsuDependencies) {
	// Grupo de rutas para los jutsus
	jutsuGroup := router.Group("/v1/jutsus")
	{
		jutsuGroup.POST("", deps.CreateJutsuController.Run)
		jutsuGroup.PUT("/:id", deps.UpdateJutsuController.Run)
		jutsuGroup.GET("", deps.GetAllJutsuController)
		jutsuGroup.DELETE("/:id", deps.DeleteJutsuController.Run)
	}
}
