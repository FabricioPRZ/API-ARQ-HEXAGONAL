package main

import (
	"API-HEXAGONAL/src/core"
	"API-HEXAGONAL/src/jutsu/dependenciesJutsu"
	"API-HEXAGONAL/src/jutsu/infrastructureJutsu"
	"API-HEXAGONAL/src/jutsu/infrastructureJutsu/adapter"
	"API-HEXAGONAL/src/jutsu/infrastructureJutsu/routesJutsu"
	"API-HEXAGONAL/src/shinobi/dependencies"
	"API-HEXAGONAL/src/shinobi/infrastructure"
	"API-HEXAGONAL/src/shinobi/infrastructure/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Conectar a la base de datos
	core.ConnectToDatabase()

	// Crear repositorios de shinobis y jutsus
	shinobiRepo := infrastructure.NewMySQLRepositoryShinobi()
	jutsuRepo := infrastructureJutsu.NewMySQLRepositoryJutsu()

	// Crear y configurar el adaptador RabbitMQ
	rabbitMQAdapter, err := adapter.NewRabbitMQAdapter()
	if err != nil {
		log.Fatal("Error al conectar con RabbitMQ:", err)
	}
	defer rabbitMQAdapter.Close()

	// Configurar dependencias de shinobis
	shinobiDeps := dependencies.NewShinobiDependencies(shinobiRepo)

	// Configurar dependencias de jutsus (incluyendo la inyección del adaptador de RabbitMQ)
	jutsusDeps := dependenciesJutsu.NewJutsuDependencies(jutsuRepo, rabbitMQAdapter)

	// Inicializar el enrutador de Gin
	router := gin.Default()

	// Configurar las rutas para shinobis
	routes.SetupShinobiRoutes(router, shinobiDeps)

	// Configurar las rutas para jutsus
	routesJutsu.SetupJutsuRoutes(router, jutsusDeps)

	// Iniciar el servidor en el puerto 8080
	log.Println("Iniciando servidor en el puerto 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
