package entity

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID         int       `gorm:"primary key; autoIncrement" db:"id" json:"id"`
	Name       string    `db:"name" gorm:"uniqueIndex" json:"name"`
	UserId     int       `db:"user_id" json:"user_id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	ClientName string    `db:"client_name" json:"client_name"`
	User       User      `gorm:"foreignKey:UserId"`
}

func MigrateProject(db *gorm.DB) error {
	return db.AutoMigrate(&Project{})
}
