package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var db *gorm.DB
var DB *gorm.DB

// Modelos
type InformacionAdmin struct {
	CodigoAdmin uint   `gorm:"primaryKey"`
	Nombre      string `gorm:"size:255"`
	Apellido    string `gorm:"size:255"`
}

type LoginAdmin struct {
	Password           string `gorm:"size:255"`
	FKInformacionAdmin uint
}

type InformacionEstudiante struct {
	CodigoEstudiante uint   `gorm:"primaryKey"`
	Nombre           string `gorm:"size:255"`
	Apellido         string `gorm:"size:255"`
}

type LoginEstudiante struct {
	Password                string `gorm:"size:255"`
	FKInformacionEstudiante uint
}

type InscripcionComedor struct {
	Contador                uint
	FKInformacionEstudiante uint
}

// InfoQR representa la información de un código QR almacenada en la base de datos
type InfoQR struct {
	FKInformacionEstudiante uint   `gorm:"primaryKey"`
	Encrypted               string `gorm:"type:text"` // La cadena encriptada
	QRRepresentation        string `gorm:"type:text"` // La representación del código QR
}

// Inicializar la conexión a la base de datos
func SetupDatabase() {
	dsn := "user=postgres password=postgres dbname=comedorgo sslmode=disable"
	var err error
	//informacion sobre lo que esta pasando en gorm
	//DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}

	// Verificar la conexión e imprimir información
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Error al obtener instancia de conexión de la base de datos: " + err.Error())
	}

	err = sqlDB.Ping()
	if err != nil {
		panic("Error al realizar ping a la base de datos: " + err.Error())
	}

	fmt.Println("Conexión a la base de datos establecida correctamente.")

	// Auto migrar modelos

	err = DB.AutoMigrate(
		&InformacionAdmin{},
		&LoginAdmin{},
		&InformacionEstudiante{},
		&LoginEstudiante{},
		&InscripcionComedor{},
		&InfoQR{},
	)
	if err != nil {
		panic("Error al migrar modelos: " + err.Error())
	}
}
