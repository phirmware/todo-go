package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"revision/todo/models"
	"revision/todo/response"
	"revision/todo/services"
	"revision/todo/utils"
	"time"

	"github.com/gorilla/mux"
)

// Todo defines the shape of the todo struct
type Todo struct {
	todoS *services.TodoService
	utils *utils.Utils
}

// NewTodo returns the new todo object
func NewTodo(todoService *services.TodoService) *Todo {
	return &Todo{
		todoS: todoService,
		utils: utils.NewUtils(),
	}
}

// GetTodos is the http handler for get todos
func (t *Todo) GetTodos(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, nil, http.StatusOK)
}

// GetATodo fetches a single todo
func (t *Todo) GetATodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	respondJSON(w, name, http.StatusOK)

}

// CreateAccount creates a todo account
func (t *Todo) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var body *models.UserModel
	json.NewDecoder(r.Body).Decode(&body)

	body.CreatedAt = time.Now()
	fmt.Printf("%+v", body)
	_, err := t.todoS.CreateUser(body)
	if err != nil {
		respondJSON(w, err, http.StatusInternalServerError)
		return
	}

	respondJSON(w, body, http.StatusOK)
}

// AddTodo adds a todo to the list
func (t *Todo) AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo *models.Todo
	t.utils.ParseForm(r, &todo)

	todo.Exp = time.Now()
	fmt.Printf("%+v", todo)

	if _, err := t.todoS.CreateTodo(todo); err != nil {
		respondJSON(w, err, http.StatusInternalServerError)
		return
	}
	respondJSON(w, todo, http.StatusOK)
}

func respondJSON(w http.ResponseWriter, data interface{}, status int) {
	res := response.NewResponse(w, status)
	res.JSON(data)
}
