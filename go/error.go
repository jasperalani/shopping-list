package main

import (
	"log"
	"net/http"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Create HTTPNotFound error response
// Need both w and r to be a valid http.HandlerFunc type
func HTTPNotFound(w http.ResponseWriter, r *http.Request) {
	createErrorResponse(w, "err_httpnotfound")
}

// Create IDNotFound error response
func IDNotFound(w http.ResponseWriter, r *http.Request) {
	createErrorResponse(w, "err_idnotfound")
}

// Create NoDataProvided error response
func NoDataProvided(w http.ResponseWriter, r *http.Request) {
	createErrorResponse(w, "err_nodataprovided")
}

// Create NoItems error response
func NoItems(w http.ResponseWriter, r *http.Request) {
	createErrorResponse(w, "err_noitems")
}
