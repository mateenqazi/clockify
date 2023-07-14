package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primary key; autoIncrement" json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}