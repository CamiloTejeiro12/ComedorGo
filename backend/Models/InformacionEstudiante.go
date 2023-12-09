package Models

type InformacionEstudiante struct {
	CodigoEstudiante uint   `gorm:"primaryKey"`
	Nombre           string `gorm:"size:255"`
	Apellido         string `gorm:"size:255"`
}
