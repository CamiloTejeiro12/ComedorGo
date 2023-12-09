package Models

type LoginEstudiante struct {
	Password                string `gorm:"size:255"`
	FKInformacionEstudiante uint
}
