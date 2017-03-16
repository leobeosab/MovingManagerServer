package main

import (
	"net/http"
)

func main() {
	router := CreateRouter()
	http.ListenAndServe(":8080", router)
}
