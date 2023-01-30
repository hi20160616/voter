package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func conn() error {
	// Configure the database connection (always check errors)
	DB, err := sql.Open("mysql", "root:rootpassword@(127.0.0.1:3306)/dbname?parseTime=true")
	if err != nil {
		return err
	}

	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	if err = DB.Ping(); err != nil {
		return err
	}
	return err
}

func init() {
	if err := conn(); err != nil {
		log.Fatal(err)
	}
}
