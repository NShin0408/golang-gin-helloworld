package models

type Recommend struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	BookID uint `gorm:"not null"`
	Book   Book `gorm:"foreignKey:BookID"`
	Detail string
	Url    string `gorm:"size:255;not null"`
}
