package main

import (
	"encoding/json"
	"net/http"
	//"strconv"
)

func createErrorResponse(w http.ResponseWriter, queryColumn string) bool {

	results := db.QueryRow("SELECT " + queryColumn + " FROM errors;")

	response := &Response{}

	results.Scan(&response.Error)

	data, _ := json.Marshal(response)

	w.Write(data)

	return true

}

func createResponse(w http.ResponseWriter, response string) {

	responseObj := &Response{
		Response: response,
	}

	data, _ := json.Marshal(responseObj)

	w.Write(data)

}

func evaluator(subject bool, outcome1 string, outcome2 string) string {
	if subject {
		return outcome1
	}
	return outcome2
}

//func funcEvaluator (subject bool, outcome1 func(Type reflect.Type))
