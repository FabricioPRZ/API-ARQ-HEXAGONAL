package main

import (
	"API-HEXAGONAL/src/core"
	"API-HEXAGONAL/src/jutsu/dependencies_jutsu"
	"API-HEXAGONAL/src/jutsu/infrastructure_jutsu"
	"API-HEXAGONAL/src/jutsu/infrastructure_jutsu/routes_jutsu"
	"API-HEXAGONAL/src/shinobi/dependencies"
	"API-HEXAGONAL/src/shinobi/infrastructure"
	"API-HEXAGONAL/src/shinobi/infrastructure/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Conectar a la base de datos
	core.ConnectToDatabase()

	// Crear repositorio de shinobis
	shinobiRepo := infrastructure.NewMySQLRepositoryShinobi()
	jutsuRepo := infrastructure_jutsu.NewMySQLRepositoryJutsu()

	// Configurar dependencias de shinobis
	shinobiDeps := dependencies.NewShinobiDependencies(shinobiRepo)
	jutsusDeps := dependencies_jutsu.NewJutsuDependencies(jutsuRepo)

	// Inicializar el enrutador
	router := gin.Default()

	// Configurar las rutas para shinobis
	routes.SetupShinobiRoutes(router, shinobiDeps)
	routes_jutsu.SetupJutsuRoutes(router, jutsusDeps)

	// Iniciar el servidor en el puerto 8080
	log.Println("Iniciando servidor en el puerto 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
