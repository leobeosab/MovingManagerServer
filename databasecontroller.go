package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//ConnectDatabase : Used to return a connection to the database
func ConnectDatabase() *sql.DB {
	instance, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/movingmanager")
	if err != nil {
		panic(err)
	}

	return instance
}

//ListMoves : Returns a slice of Moves
func ListMoves() *[]Move {
	db := ConnectDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM moves")
	if err != nil {
		fmt.Println(err.Error())
	}

	data := make([]Move, 0, 0)
	var i, o, d, f, p []byte

	for rows.Next() {
		err = rows.Scan(&i, &o, &d, &f, &p)
		if err != nil {
			fmt.Println(err.Error())
		}

		data = append(data, Move{ID: string(i), OldAddress: string(o), DestinationAddress: string(d), FamilyName: string(f), PreviewImageURL: string(p)})
	}

	return &data
}

//AddMoveToDB : Adds a Move struct to the MySQL database
func AddMoveToDB(move *Move) string {
	db := ConnectDatabase()
	defer db.Close()

	status := "success"

	_, err := db.Query("INSERT INTO moves VALUES (0, \"" + move.OldAddress + "\", \"" + move.DestinationAddress + "\", \"" + move.FamilyName + "\", \"" + move.PreviewImageURL + "\")")
	if err != nil {
		status = "failure"
		fmt.Println(err.Error())
	}

	return status
}

//RemoveMoveFromDB : Removes a Move from the database
func RemoveMoveFromDB(id string) string {
	db := ConnectDatabase()
	defer db.Close()

	status := "success"

	_, err := db.Query("DELETE FROM moves WHERE id = " + id)
	if err != nil {
		status = "failure"
		fmt.Println(err.Error())
	}

	return status
}

//Helpers:

//GetStringFromDB : Used for getting Ids mostly, DB connection in so we don't open more than we need
func GetStringFromDB(query string, db *sql.DB) string {

	row, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	if !row.Next() {
		return ""
	}

	var tID []byte
	row.Scan(&tID)

	return string(tID)
}
