package models

import "time"

// UserModel defines the model of a database user
type UserModel struct {
	ID        uint
	Username  string
	Password  string
	CreatedAt time.Time
}
