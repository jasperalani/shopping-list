package main

import (
	"database/sql"
	"encoding/json"
	//"log"
	"net/http"
	//"strconv"
)

func createErrorResponse(w http.ResponseWriter, queryColumn string) bool {

	results := db.QueryRow("SELECT " + queryColumn + " FROM errors;")

	response := &Response{
		Errno: 1,
	}

	err = results.Scan(&response.Error)
	handleError(err)

	data, err := json.Marshal(response)
	handleError(err)

	_, err = w.Write(data)
	handleError(err)

	return true

}

func createResponse(w http.ResponseWriter, response string) {

	responseObj := &Response{
		Response: response,
		Errno:    0,
	}

	data, err := json.Marshal(responseObj)
	handleError(err)

	_, err = w.Write(data)
	handleError(err)

}

func evaluator(subject bool, outcome1 string, outcome2 string) string {
	if subject {
		return outcome1
	}
	return outcome2
}

func selectID(id string) float32 {

	/*

		Error codes:
		-11 sql: no rows in result set

	*/

	var (
		itemExists bool
		queryID    string = "SELECT id FROM items WHERE id = " + id
		item       Item
	)

	err = db.Get(&item, queryID)

	if err == sql.ErrNoRows {
		return 0
	} else {
		handleError(err)
	}

	itemExists = item.ID != 0

	if !itemExists {
		return 0
	}

	return float32(item.ID)

}

func anyItems() bool {

	var (
		items    []Item
		queryAll string = "SELECT * FROM `shopping-list`.items;"
	)

	err = db.Select(&items, queryAll)
	return err != sql.ErrNoRows

}

//func funcEvaluator (subject bool, outcome1 func(Type reflect.Type))
