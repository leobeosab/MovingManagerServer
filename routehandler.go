//Route handling definitions
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Index : Used to return a list of Moves
func Index(w http.ResponseWriter, r *http.Request) {

	moves := MoveResponse{Moves: *ListMoves()}

	marshalled, err := json.Marshal(&moves)
	if err != nil {
		fmt.Println(err.Error())
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "{ \"status\": \"Error retrieving moves from database\" }")
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(marshalled))
	}
}

//AddMove : Adds move to the MySQL database
func AddMove(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hah")

	status := "success"

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var m Move

	err := decoder.Decode(&m)
	if err != nil {
		status = "failure"
		log.Fatal(err)
	}

	if status == "success" {
		status = AddMoveToDB(&m)
	}

	resp, err := json.Marshal(&Response{Status: status})
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(resp))
}

//DeleteMove : Removes move from the MySQL database
func DeleteMove(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	status := RemoveMoveFromDB(id)

	resp, err := json.Marshal(&Response{Status: status})
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(resp))
}
