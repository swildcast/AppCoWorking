package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas las rutas del backend
func SetupRoutes(r *gin.Engine) {
	// Grupo de rutas de autenticación
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}

	// Grupo de rutas de logs de actividad
	logRoutes := r.Group("/api/logs")
	{
		logRoutes.POST("/", controllers.CreateLog)
		logRoutes.GET("/", controllers.GetLogs)
	}

	// Grupo de rutas protegidas con autenticación
	protectedRoutes := r.Group("/api/protected").Use(controllers.AuthMiddleware())
	{
		protectedRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Ruta protegida"})
		})
	}
}
