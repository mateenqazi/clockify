package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primary key; autoIncrement" db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	IsActive  bool      `db:"is_active" json:"is_active"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
