package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	//"strconv"
)

func createErrorResponse(w http.ResponseWriter, queryColumn string) {

	db, _ := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/shoppinglist")

	results := db.QueryRow("SELECT " + queryColumn + " FROM errors")

	response := &Response{}

	results.Scan(&response.Error)

	data, _ := json.Marshal(response)

	w.Write(data)

}

func createResponse(w http.ResponseWriter, response string) {

	responseObj := &Response{
		Response: response,
	}

	data, _ := json.Marshal(responseObj)

	w.Write(data)

}
