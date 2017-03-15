//Route handling definitions

package main

import (
	"fmt"
	"net/http"
)

//Print out the version
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Notification Server API: 0.1.0")
}

//Subscribe user based on available data and recieved data
func AddMove(w http.ResponseWriter, r *http.Request) {

}

func DeleteMove(w http.ResponseWriter, r *http.Request) {

}
