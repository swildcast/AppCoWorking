package seeders

import (
	"backend/config"
	"backend/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func RunSeeders() {
	// Crear roles
	roles := []models.Role{
		{Nombre: "admin"},
		{Nombre: "cliente"},
		{Nombre: "recepcionista"},
	}

	for _, role := range roles {
		config.DB.Create(&role)
	}

	// Crear un usuario admin
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	adminUser := models.User{
		Nombre:   "Admin",
		Email:    "admin@example.com",
		Password: string(hashedPassword),
		RolID:    1, // ID del rol "admin"
	}

	config.DB.Create(&adminUser)

	log.Println("✅ Seeders ejecutados correctamente")
}
