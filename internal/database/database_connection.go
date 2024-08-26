package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func GetConnection() *sql.DB {
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	url := os.Getenv("DATABASE_URL")
	port := os.Getenv("DATABASE_PORT")
	schema := os.Getenv("DATABASE_SCHEMA")
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+url+":"+port+")/"+schema)
	if err != nil {
		log.Panicf(err.Error())
	}
	return db
}
