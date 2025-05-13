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

// Conexión a la base de datos y migraciones
func ConnectDatabase() {
	// Construcción de la cadena de conexión desde variables de entorno
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Conectar con PostgreSQL
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Conexión exitosa a PostgreSQL")

	// Migrar las tablas (solo crea las que no existan)
	database.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Space{},
		&models.Booking{},
		&models.Payment{},
		&models.ActivityLog{},
	)

	fmt.Println("Migraciones aplicadas correctamente")

	// Asignar la conexión global
	DB = database

	// Insertar datos de prueba
	SeedDatabase()
}

// Insertar datos iniciales
func SeedDatabase() {
	fmt.Println("Insertando datos iniciales...")

	// Crear roles
	roles := []models.Role{
		{Nombre: "Administrador"},
		{Nombre: "Usuario"},
	}
	if result := DB.Create(&roles); result.Error != nil {
		fmt.Println("Error insertando roles:", result.Error)
	}

	// Crear usuarios
	users := []models.User{
		{Nombre: "Admin", Email: "admin@example.com", Password: "123456", RolID: 1, FechaRegistro: time.Now()},
		{Nombre: "Usuario1", Email: "user1@example.com", Password: "password", RolID: 2, FechaRegistro: time.Now()},
	}
	if result := DB.Create(&users); result.Error != nil {
		fmt.Println("Error insertando usuarios:", result.Error)
	}

	// Crear espacios (actualizado con todos los campos requeridos)
	spaces := []models.Space{
		{
			Nombre:        "Sala de Reuniones A",
			Descripcion:   "Sala equipada para 10 personas",
			Capacidad:     10,
			PrecioPorHora: 50.00,
			Estado:        "disponible",
			Ubicacion:     "Piso 1",
		},
		{
			Nombre:        "Sala de Conferencias B",
			Descripcion:   "Amplio espacio para eventos",
			Capacidad:     50,
			PrecioPorHora: 120.00,
			Estado:        "disponible",
			Ubicacion:     "Piso 2",
		},
	}
	if result := DB.Create(&spaces); result.Error != nil {
		fmt.Println("Error insertando espacios:", result.Error)
	}

	// Crear reservas
	bookings := []models.Booking{
		{UsuarioID: 1, EspacioID: 1, FechaInicio: time.Now().AddDate(0, 0, 1), FechaFin: time.Now().AddDate(0, 0, 1).Add(2 * time.Hour), Estado: "confirmada"},
		{UsuarioID: 2, EspacioID: 2, FechaInicio: time.Now().AddDate(0, 0, 2), FechaFin: time.Now().AddDate(0, 0, 2).Add(3 * time.Hour), Estado: "pendiente"},
	}
	if result := DB.Create(&bookings); result.Error != nil {
		fmt.Println("Error insertando reservas:", result.Error)
	}

	// Crear pagos basados en las reservas existentes
	var reservas []models.Booking
	DB.Find(&reservas)

	if len(reservas) >= 2 {
		payments := []models.Payment{
			{ReservationID: reservas[0].ID, Amount: 100.50, PaymentMethod: "Tarjeta", Status: "Pagado", CreatedAt: time.Now()},
			{ReservationID: reservas[1].ID, Amount: 50.00, PaymentMethod: "Efectivo", Status: "Pendiente", CreatedAt: time.Now()},
		}
		if result := DB.Create(&payments); result.Error != nil {
			fmt.Println("Error insertando pagos:", result.Error)
		}
	}

	// Crear logs de actividad
	logs := []models.ActivityLog{
		{UserID: 1, Accion: "Inicio de sesión", Detalles: "Usuario Admin inició sesión"},
		{UserID: 2, Accion: "Reserva creada", Detalles: "Usuario1 reservó la Sala de Conferencias B"},
	}
	if result := DB.Create(&logs); result.Error != nil {
		fmt.Println("Error insertando logs:", result.Error)
	}

	fmt.Println("Datos iniciales insertados correctamente")
}
