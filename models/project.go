package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID         int       `gorm:"primary key; autoIncrement" db:"id"`
	Name       string    `db:"name" gorm:"uniqueIndex"`
	UserId     string    `db:"user_id"`
	CreatedAt  time.Time `db:"created_at"`
	ClientName string    `db:"client_name"`
	User       User      `gorm:"foreignKey:UserId"`
}

func MigrateProject(db *gorm.DB) error {
	return db.AutoMigrate(&Project{})
}
