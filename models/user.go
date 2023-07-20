package models

import (
	"time"

	"gorm.io/gorm"
)

type UserInterface interface {
	IsEmailExists(db *gorm.DB, email string) (bool, User, error)
}

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

func (u *User) IsEmailExists(db *gorm.DB, email string) (bool, User, error) {
	var user User

	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, user, nil
		}

		return false, user, result.Error
	}

	return true, user, nil
}
