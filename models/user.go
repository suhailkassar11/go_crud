package models

import "gorm.io/gorm"

type User struct {
	Username string `json:"username" gorm:"not null" validate:"required`
	Email    string `json:"email" gorm:"uniqueIndex;not null" validate:required"`
	Password string `json:"password" gorm:"not null" validate:"required"`
	gorm.Model
}
