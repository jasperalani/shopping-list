package main

import "net/http"

// Create HTTPNotFound error response
func HTTPNotFound(w http.ResponseWriter) {
	createErrorResponse(w, "err_httpnotfound")
}

// Create IDNotFound error response
func IDNotFound(w http.ResponseWriter) {
	createErrorResponse(w, "err_idnotfound")
}

// Create NoDataProvided error response
func NoDataProvided(w http.ResponseWriter) {
	createErrorResponse(w, "err_nodataprovided")
}

// Create NoItems error response
func NoItems(w http.ResponseWriter) {
	createErrorResponse(w, "err_noitems")
}
