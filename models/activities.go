package models

import (
	"time"

	"gorm.io/gorm"
)

type Activities struct {
	ID           int       `gorm:"primary key; autoIncrement" json:"id"`
	Name         string    `json:"name"`
	TimeDuration time.Time `json:"timeDuration"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	ProjectId    int       `json:"projectId"`
	UserId       int       `json:"userId"`
	User         User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Project      Project   `gorm:"foreignKey:ProjectId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func MigrateActivities(db *gorm.DB) error {
	err := db.AutoMigrate(&Activities{})
	return err
}
