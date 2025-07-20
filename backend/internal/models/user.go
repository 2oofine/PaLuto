package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"notnu"`
}
