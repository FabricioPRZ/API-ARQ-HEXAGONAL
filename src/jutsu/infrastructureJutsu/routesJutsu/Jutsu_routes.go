package routesJutsu

import (
	"API-HEXAGONAL/src/jutsu/dependenciesJutsu"
	"github.com/gin-gonic/gin"
)

// Configura las rutas para los jutsus
func SetupJutsuRoutes(router *gin.Engine, deps *dependenciesJutsu.JutsuDependencies) {
	// Grupo de rutas para los jutsus
	jutsuGroup := router.Group("/v1/jutsus")
	{
		jutsuGroup.POST("", deps.CreateJutsuController.Run)
		jutsuGroup.PUT("/:id", deps.UpdateJutsuController.Run)
		jutsuGroup.GET("", deps.GetAllJutsuController)
		jutsuGroup.GET("/:id", deps.GetByIdJutsuController)
		jutsuGroup.DELETE("/:id", deps.DeleteJutsuController.Run)
	}
}
