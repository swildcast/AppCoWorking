package main

import (
	"backend/config"
	"backend/routes"
	"backend/seeders"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Conectar a la base de datos
	config.ConnectDatabase()

	// Ejecutar seeders para datos iniciales
	seeders.RunSeeders()

	// Inicializar el router de Gin
	r := gin.Default()

	// Configurar las rutas
	routes.SetupRoutes(r)

	// Iniciar el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
