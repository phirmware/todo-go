package main

import (
	"net/http"
	"revision/todo/adapters"
	"revision/todo/routes"
	_ "github.com/go-sql-driver/mysql"
)

const (
	port = ":8080"
)

func main() {

	db := adapters.ConnectDB()
	r := routes.Init(db)

	http.ListenAndServe(port, r)
}
