package types

import "time"

type Credentials struct {
	Email    string
	Password string
}

type User struct {
	ID        int
	Email     string
	Password  string
	CreatedAt time.Time
	IsActive  bool
}
