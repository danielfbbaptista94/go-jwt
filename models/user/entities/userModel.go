package entities

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}
