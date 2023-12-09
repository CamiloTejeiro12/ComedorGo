package Models

type LoginAdmin struct {
	Password           string `gorm:"size:255"`
	FKInformacionAdmin uint
}
