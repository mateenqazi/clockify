package entity

import (
	"time"

	"gorm.io/gorm"
)

type Activities struct {
	ID           int           `gorm:"primary key; autoIncrement" db:"id" json:"id"`
	Name         string        `db:"name" gorm:"uniqueIndex" json:"name"`
	TimeDuration time.Duration `db:"time_duration" json:"time_duration"`
	StartTime    time.Time     `db:"start_time" json:"start_time"`
	EndTime      time.Time     `db:"end_time" json:"end_time"`
	ProjectId    int           `db:"project_id" json:"project_id"`
	UserId       int           `db:"user_id" json:"user_id"`
	User         User          `gorm:"foreignKey:UserId"`
	Project      Project       `gorm:"foreignKey:ProjectId"`
}

func MigrateActivities(db *gorm.DB) error {
	return db.AutoMigrate(&Activities{})
}
