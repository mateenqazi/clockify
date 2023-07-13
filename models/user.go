package models

import "time"

type User struct {
	id         int
	email      string
	password   string
	created_at time.Time
	is_active  bool
}
