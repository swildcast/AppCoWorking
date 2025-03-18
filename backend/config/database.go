package config

import (
	"backend/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variable global de la base de datos
var DB *gorm.DB

// 📌 Conexión a la base de datos y migraciones
func ConnectDB() {
	// Construcción de la cadena de conexión desde variables de entorno
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Conectar con PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error al conectar a la base de datos:", err)
	}

	fmt.Println("✅ Conexión exitosa a PostgreSQL")

	// 🔥 Eliminar tablas antes de migrar (⚠️ Esto borra TODOS los datos)
	db.Migrator().DropTable(
		&models.User{},
		&models.Booking{},
		&models.Payment{},
		&models.ActivityLog{},
	)

	// 🔄 Migrar las tablas
	db.AutoMigrate(
		&models.User{},
		&models.Booking{},
		&models.Payment{},
		&models.ActivityLog{},
	)

	fmt.Println("🚀 Migraciones aplicadas correctamente.")

	// Asignar la conexión global
	DB = db

	// 🛠️ Insertar datos de prueba
	SeedDatabase()
}

// 📌 Insertar datos iniciales
func SeedDatabase() {
	// 🔹 Crear roles
	roles := []models.Role{
		{Nombre: "Administrador"},
		{Nombre: "Usuario"},
	}
	DB.Create(&roles)

	// 🔹 Crear usuarios
	users := []models.User{
		{Nombre: "Admin", Email: "admin@example.com", Password: "123456", RolID: 1, FechaRegistro: time.Now()},
		{Nombre: "Usuario1", Email: "user1@example.com", Password: "password", RolID: 2, FechaRegistro: time.Now()},
	}
	DB.Create(&users)

	// 🔹 Crear reservas
	bookings := []models.Booking{
		{UsuarioID: 1, EspacioID: 1, FechaInicio: time.Now().AddDate(0, 0, 1), FechaFin: time.Now().AddDate(0, 0, 1).Add(2 * time.Hour), Estado: "confirmada"},
		{UsuarioID: 2, EspacioID: 2, FechaInicio: time.Now().AddDate(0, 0, 2), FechaFin: time.Now().AddDate(0, 0, 2).Add(3 * time.Hour), Estado: "pendiente"},
	}
	DB.Create(&bookings)

	// 🔹 Crear pagos basados en las reservas existentes
	var reservas []models.Booking
	DB.Find(&reservas)

	if len(reservas) >= 2 {
		payments := []models.Payment{
			{ReservationID: reservas[0].ID, Amount: 100.50, PaymentMethod: "Tarjeta", Status: "Pagado", CreatedAt: time.Now()},
			{ReservationID: reservas[1].ID, Amount: 50.00, PaymentMethod: "Efectivo", Status: "Pendiente", CreatedAt: time.Now()},
		}
		DB.Create(&payments)
	}

	// 🔹 Crear logs de actividad
	logs := []models.ActivityLog{
		{UserID: 1, Accion: "Inicio de sesión", Detalles: "Usuario Admin inició sesión"},
		{UserID: 2, Accion: "Reserva creada", Detalles: "Usuario1 reservó la Sala de Conferencias B"},
	}
	DB.Create(&logs)

	fmt.Println("✅ Datos iniciales insertados correctamente.")
}
