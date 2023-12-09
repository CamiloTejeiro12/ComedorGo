package Models

type InfoQR struct {
	FKInformacionEstudiante uint   `gorm:"primaryKey"`
	Encrypted               string `gorm:"type:text"` // La cadena encriptada
	QRRepresentation        string `gorm:"type:text"` // La representación del código QR
}
