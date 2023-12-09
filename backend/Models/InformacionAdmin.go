package Models

type InformacionAdmin struct {
	CodigoAdmin uint   `gorm:"primaryKey"`
	Nombre      string `gorm:"size:255"`
	Apellido    string `gorm:"size:255"`
}
