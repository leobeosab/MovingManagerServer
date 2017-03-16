//Route data
package main

import "net/http"

//Route : Used for handling server requests
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Response : Used for json feedback
type Response struct {
	Status string `json:"status"`
}

//Move : Contianer to hold data stored in MySQL db.
type Move struct {
	ID                 string `json:"id"`
	OldAddress         string `json:"old_address"`
	DestinationAddress string `json:"destination_address"`
	FamilyName         string `json:"family_name"`
	PreviewImageURL    string `json:"preview_image_url"`
}
