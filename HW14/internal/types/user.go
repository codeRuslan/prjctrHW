package types

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name       string `gorm:"not null"`
	SecondName string `gorm:"not null"`
}
