package models

import (
	"time"

	"gorm.io/gorm"
)

type Activities struct {
	ID           int       `gorm:"primary key; autoIncrement" db:"id"`
	Name         string    `db:"name"`
	TimeDuration time.Time `db:"time_duration"`
	StartTime    time.Time `db:"start_time"`
	EndTime      time.Time `db:"end_time"`
	ProjectId    int       `db:"project_id"`
	UserId       int       `db:"user_id"`
	User         User      `gorm:"foreignKey:UserId"`
	Project      Project   `gorm:"foreignKey:ProjectId"`
}

func MigrateActivities(db *gorm.DB) error {
	return db.AutoMigrate(&Activities{})
}
