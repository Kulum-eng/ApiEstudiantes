package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"API-HEXAGONAL/src/core"
	"API-HEXAGONAL/src/jutsu/dependenciesJutsu"
	"API-HEXAGONAL/src/jutsu/infrastructureJutsu"
	"API-HEXAGONAL/src/jutsu/infrastructureJutsu/routesJutsu"
	"API-HEXAGONAL/src/estudiantes/dependencies"
	"API-HEXAGONAL/src/estudiantes/infrastructure"
	"API-HEXAGONAL/src/estudiantes/infrastructure/routes"
)

func main() {
	// Conectar a la base de datos
	core.ConnectToDatabase()

	// Crear repositorios de shinobis y jutsus
	shinobiRepo := infrastructure.NewMySQLRepositoryShinobi()
	jutsuRepo := infrastructureJutsu.NewMySQLRepositoryJutsu()

	// Configurar dependencias de shinobis
	shinobiDeps := dependencies.NewShinobiDependencies(shinobiRepo)

	jutsusDeps := dependenciesJutsu.NewJutsuDependencies(jutsuRepo)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))


	routes.SetupShinobiRoutes(router, shinobiDeps)


	routesJutsu.SetupJutsuRoutes(router, jutsusDeps)

	
	log.Println("Iniciando servidor en el puerto 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
