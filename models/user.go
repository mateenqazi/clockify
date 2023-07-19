package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primary key; autoIncrement" db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	IsActive  bool      `db:"is_active"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
