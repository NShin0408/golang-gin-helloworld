package models

import (
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(100);unique_index"`
}
