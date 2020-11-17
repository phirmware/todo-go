package routes

import (
	"database/sql"
	"revision/todo/controllers"
	"revision/todo/services"

	"github.com/gorilla/mux"
)

// Init initializes the router object
func Init(db *sql.DB) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	todoS := services.NewTodoService(db)
	todoC := controllers.NewTodo(todoS)

	r.HandleFunc("/api/users/add-new", todoC.CreateAccount).Methods("POST")

	r.HandleFunc("/api/todos", todoC.GetTodos).Methods("GET")
	r.HandleFunc("/api/todos/add-new", todoC.AddTodo).Methods("POST")

	return r
}
