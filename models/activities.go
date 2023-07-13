package models

import "time"

type Activities struct {
	id            int
	name          string
	time_duration time.Time
	start_time    time.Time
	end_time      time.Time
	project_id    int
	user_id       int
}
