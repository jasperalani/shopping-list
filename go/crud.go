package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

var (
	databaseUrl string = "root:password@tcp(127.0.0.1:3306)/shopping-list"
	db, _              = sql.Open("mysql", databaseUrl)
)

func createItemRecord(w http.ResponseWriter, request *http.Request) {

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

		createResponse(w, []string{"inc_qty", "Increased quantity of matching item, did not insert item"})

	} else {

		insertQuery := "INSERT INTO items (name, url, image_url, person, quantity) "
		insertQuery = insertQuery + "VALUES ('" + item.Name + "', '" + item.URL + "', '" + item.ImageURL + "', '" + item.Person + "', " + strconv.Itoa(item.Quantity) + ")"

		db.Query(insertQuery)

		json.NewEncoder(w).Encode(&item)

	}

}

func readItemRecord(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	if len(params) > 0 { //If number
		db, _ := sql.Open("mysql", databaseUrl)

		var id int

		selectIDQuery := db.QueryRow("SELECT id FROM items WHERE id = " + params["id"])

		selectIDQuery.Scan(&id)

		if id == 0 {

			createErrorResponse(w, "err_idnotfound")

		} else {

			var (
				item  []Item
				query string = "SELECT id, name, url, image_url, person FROM items WHERE id = " + params["id"]
			)

			db, err := sqlx.Connect("mysql", "root:password@tcp(127.0.0.1:3306)/shopping-list")
			returnError(err)

			err = db.Select(&item, query)
			returnError(err)

			err = json.NewEncoder(w).Encode(&item)
			returnError(err)

			//results, _ := db.Query()
			//
			//for results.Next() {
			//
			//
			//	//results.Scan(
			//	//	&item.ID,
			//	//	&item.Name,
			//	//	&item.URL,
			//	//	&item.ImageURL,
			//	//	&item.Person,
			//	//	&item.Quantity,
			//	//	&item.Deleted,
			//	//	)
			//
			//	results.Scan(
			//		&item,
			//	)
			//
			//	json.NewEncoder(w).Encode(&item)
			//
			//}

		}

	} else { // if all

		db, _ := sql.Open("mysql", databaseUrl)

		results, err := db.Query("SELECT * FROM items")

		if err != nil {
			log.Panic(err)
		}

		var items []Item

		for results.Next() {

			var item Item

			results.Scan(&item.ID, &item.Name, &item.URL, &item.ImageURL, &item.Person, &item.Quantity)

			if reflect.TypeOf(item.ID) != nil {
				items = append(items, item)
			}
		}

		if len(items) > 0 {
			json.NewEncoder(w).Encode(&items)
		} else {
			NoItems(w, r)
		}

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
