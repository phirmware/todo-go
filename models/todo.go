package models

import "time"

// Todo defines the model of the todo data
type Todo struct {
	Name        string `json:"name"`
	UserID      int    `json:"userId"`
	Description string `json:"description"`
	Exp         time.Time
}
