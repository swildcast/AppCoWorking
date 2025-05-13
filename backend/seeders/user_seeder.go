package seeders

import (
	"backend/config"
	"backend/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func SeedUsers() {
	password, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error hashing password:", err)
	}

	users := []models.User{
		{
			Nombre:   "Usuario Prueba",
			Email:    "test@example.com",
			Password: string(password),
		},
	}

	for _, user := range users {
		if err := config.DB.Create(&user).Error; err != nil {
			log.Printf("Error creating user %s: %v", user.Email, err)
		}
	}
}
