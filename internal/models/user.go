package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"uniqueIndex"`
	// хеш пароля
	PasswordHash string  `json:"-"`
	Role         string  `json:"role" gorm:"default:user"` // user | admin
	Orders       []Order `gorm:"foreignKey:UserID" json:"orders,omitempty"`
}
