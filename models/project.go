package models

import "time"

type Project struct {
	Id          int
	name        string
	user_id     string
	created_at  time.Time
	client_name string
}
