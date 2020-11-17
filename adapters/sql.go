package adapters

import (
	"database/sql"
	"fmt"
)

// ConnectDB connects to the database
func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/?parseTime=true")
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	RunMigrations(db)
	return db
}

// RunMigrations runs migrations for the DB
func RunMigrations(db *sql.DB) {
	if _, err := db.Exec(`CREATE DATABASE todo`); err != nil {
		fmt.Println(err)
	}

	if _, err := db.Exec(`USE todo`); err != nil {
		panic(err)
	}

	var userquery = `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
	);`

	var todoquery = `
    CREATE TABLE todos (
		id INT AUTO_INCREMENT,
		userid INT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        exp DATETIME,
        PRIMARY KEY (id)
	);`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	if  _, err := db.Exec(userquery);err != nil {
		fmt.Println((err))
	}

	_, err := db.Exec(todoquery)
	if err != nil {
		fmt.Println((err))
	}
}
