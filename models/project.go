package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID         int       `gorm:"primary key; autoIncrement" json:"id"`
	Name       string    `json:"name" gorm:"uniqueIndex"`
	UserId     int       `json:"userId"`
	CreatedAt  time.Time `json:"createdAt"`
	ClientName string    `json:"clientName"`
	User       User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func MigrateProject(db *gorm.DB) error {
	err := db.AutoMigrate(&Project{})
	return err
}
