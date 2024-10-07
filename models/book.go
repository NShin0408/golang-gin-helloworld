package models

type Book struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"size:100;not null"`
	Detail        string
	ImageFilename string `gorm:"not null"`
	AmazonUrl     string `gorm:"size:255;not null"`
}
