package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	instance, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/movingmanager")
	if err != nil {
		panic(err)
	}

	return instance
}

//Helpers:

//Used for getting Ids mostly, DB connection in so we don't open more than we need
func GetStringFromDB(query string, db *sql.DB) string {

	row, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	if !row.Next() {
		return ""
	}

	var tId []byte
	row.Scan(&tId)

	return string(tId)
}
