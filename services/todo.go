package services

import (
	"database/sql"
	"errors"
	"revision/todo/models"
)

// TodoService defines the shape of a todo
type TodoService struct {
	db *sql.DB
}

// NewTodoService returns a new instance of the service
func NewTodoService(db *sql.DB) *TodoService {
	return &TodoService{
		db: db,
	}
}

//CreateUser adds a user to the database
func (ts *TodoService) CreateUser(user *models.UserModel) (sql.Result, error) {
	if user.Username == "" {
		return nil, errors.New("Username missing")
	}
	return ts.db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, user.Username, user.Password, user.CreatedAt)
}
