package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	//"reflect"
	"strconv"
)

var (
	databaseUrl string = "root:password@tcp(127.0.0.1:3306)/shopping-list"
	db, err            = sqlx.Connect("mysql", "root:password@tcp(127.0.0.1:3306)/shopping-list")
)

func createItemRecord(w http.ResponseWriter, request *http.Request) {

	returnError(err) // line 16

	var item Item

	_ = json.NewDecoder(request.Body).Decode(&item)

	checkQuery := "SELECT id, name, quantity FROM items WHERE name LIKE '" + item.Name + "';"

	results := db.QueryRow(checkQuery)

	var (
		id       int
		name     string
		quantity int
	)

	results.Scan(&id, &name, &quantity)

	if item.Name == name {

		finalQuantity := item.Quantity + quantity

		updateQuery := "UPDATE items set quantity = " + strconv.Itoa(finalQuantity) + " WHERE id = " + strconv.Itoa(id)

		db.Query(updateQuery)

		createResponse(w, "inc_qty")

	} else {

		insertQuery := "INSERT INTO items (name, url, image_url, person, quantity)"
		insertQuery = insertQuery + "VALUES ('" + item.Name + "', '" + item.URL + "', '" + item.ImageURL + "', '" + item.Person + "', " + strconv.Itoa(item.Quantity) + ")"

		_, err = db.Query(insertQuery)
		returnError(err)

		var (
			maxID *sql.Rows
			ID    int
		)

		maxID, err = db.Query("SELECT MAX(id) FROM items") // this query is not safe
		returnError(err)

		if maxID.Next() {
			err = maxID.Scan(&ID)
			returnError(err)
		}

		returnedItem := &Item{
			ID:       ID,
			Name:     item.Name,
			URL:      item.URL,
			Person:   item.Person,
			Quantity: item.Quantity,
		}

		json.NewEncoder(w).Encode(returnedItem)

	}

}

func readItemRecord(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var (
		item       Item
		items      []Item
		query      string
		queryScope bool = len(params) > 0
	)

	query = evaluator(queryScope,
		"SELECT id, name, url, image_url, person, quantity, deleted FROM items WHERE id = "+params["id"],
		"SELECT * FROM items",
	)

	if queryScope {

		err = db.Get(&item, query)

		if item.ID == 0 {
			createErrorResponse(w, "err_idnotfound")
			return
		}

		json.NewEncoder(w).Encode(item)

	} else {
		err = db.Select(&items, query)

		if len(items) == 0 {
			createErrorResponse(w, "err_noitems")
			return
		}

		json.NewEncoder(w).Encode(items)
	}

}

func updateItemRecord(w http.ResponseWriter, r *http.Request) {

	/*
		{
			"name": "Marcel",
			"quantity": 6504
		}
		{
			"name": string,
			"quantity": int
		}
	*/

	params := mux.Vars(r)

	db, _ := sql.Open("mysql", "root:password@tcp(mariadb:3306)/shoppinglist")

	var item Item

	_ = json.NewDecoder(r.Body).Decode(&item)

	if len(item.Name) > 0 || len(item.URL) > 0 || len(item.ImageURL) > 0 || len(item.Person) > 0 || item.Quantity > 0 { //exists
		var (
			updateQuery string
		)
		updateQuery = "UPDATE items set "
		if len(item.Name) > 0 {
			updateQuery = updateQuery + "name = '" + item.Name + "', "
		}
		if len(item.Name) > 0 {
			updateQuery = updateQuery + "url = '" + item.URL + "', "
		}
		if len(item.Name) > 0 {
			updateQuery = updateQuery + "image_url = '" + item.ImageURL + "', "
		}
		if len(item.Name) > 0 {
			updateQuery = updateQuery + "person = '" + item.Person + "', "
		}
		if len(item.Name) > 0 {
			updateQuery = updateQuery + "quantity = " + strconv.Itoa(item.Quantity) + " "
		}
		updateQuery = updateQuery + " WHERE id = " + params["id"]

		db.Query(updateQuery)

		selectQuery := "SELECT * FROM items WHERE id = " + params["id"]

		var updatedItem Item

		db.QueryRow(selectQuery).Scan(&updatedItem.ID, &updatedItem.Name, &updatedItem.URL, &updatedItem.ImageURL, &updatedItem.Person, &updatedItem.Quantity)

		json.NewEncoder(w).Encode(&updatedItem)

	} else {
		NoDataProvided(w, r)
	}

}

func deleteItemRecord(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r) // /api/items/update/{id}

	db, _ := sql.Open("mysql", "root:password@tcp(mariadb:3306)/shoppinglist")

	_, err := db.Query("DELETE FROM items WHERE id = " + params["id"])
	if err != nil {
		log.Panic(err)
	}

}
