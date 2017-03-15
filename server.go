package main

import (
	"net/http"
)

func main() {
	router := CreateRouter()
	http.ListenAndServe(":80", router)
}
