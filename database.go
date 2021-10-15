package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getDB() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}