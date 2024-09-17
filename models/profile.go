package models

type Profile struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"size:100;not null"`
	Detail        string
	ImageFilename string `gorm:"not null"`
}
