package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/iancoleman/strcase"
	_ "github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

var (
	databaseUrl string = "root:password@tcp(127.0.0.1:3306)/shopping-list"
	db, err            = sqlx.Connect("mysql", databaseUrl)
)

func createItemRecord(w http.ResponseWriter, request *http.Request) {

	var (
		item     Item
		id       int
		name     string
		quantity int
	)

	_ = json.NewDecoder(request.Body).Decode(&item)

	checkQuery := "SELECT id, name, quantity FROM items WHERE name LIKE '" + item.Name + "';"

	results := db.QueryRow(checkQuery)

	results.Scan(&id, &name, &quantity)

	if item.Name == name {

		finalQuantity := item.Quantity + quantity

		updateQuery := "UPDATE items set quantity = " + strconv.Itoa(finalQuantity) + " WHERE id = " + strconv.Itoa(id)

		db.Query(updateQuery)

		createResponse(w, "quantity_increased")

	} else {

		insertQuery := "INSERT INTO items (name, url, image_url, person, quantity)"
		insertQuery = insertQuery + "VALUES ('" + item.Name + "', '" + item.URL + "', '" + item.ImageURL + "', '" + item.Person + "', " + strconv.Itoa(item.Quantity) + ")"

		_, err = db.Query(insertQuery)
		handleError(err)

		var (
			maxID *sql.Rows
			ID    int
		)

		maxID, err = db.Query("SELECT MAX(id) FROM items") // this query is not safe
		handleError(err)

		if maxID.Next() {
			err = maxID.Scan(&ID)
			handleError(err)
		}

		createResponse(w, "item_created")

	}

}

func readItemRecord(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var (
		item       Item
		item_      ItemJSON
		items      []Item
		items_     []ItemJSON
		query      string
		queryScope bool = len(params) > 0
	)

	query = evaluator(queryScope,
		"SELECT id, name, url, image_url, person, quantity, deleted FROM items WHERE id = "+params["id"]+";",
		"SELECT * FROM items;",
	)

	if queryScope {

		err = db.Get(&item, query)

		if item.ID == 0 {
			IDNotFound(w, r)
			return
		}

		item_ = ItemJSON{
			ID:       item.ID,
			Name:     item.Name,
			URL:      item.URL,
			ImageURL: item.ImageURL,
			Person:   item.Person,
			Quantity: item.Quantity,
			Deleted:  item.Deleted,
		}

		json.NewEncoder(w).Encode(item_)

	} else {

		if !anyItems() {
			NoItems(w, r)
			return
		}

		err = db.Select(&items, query)

		for _, item := range items {
			items_ = append(items_, ItemJSON{
				ID:       item.ID,
				Name:     item.Name,
				URL:      item.URL,
				ImageURL: item.ImageURL,
				Person:   item.Person,
				Quantity: item.Quantity,
				Deleted:  item.Deleted,
			})
		}

		err = json.NewEncoder(w).Encode(items_)
		handleError(err)
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

	var (
		id          float32
		item        ItemJSON
		updateQuery string
		fieldName   string
		values      []interface{}
		valueTypes  reflect.Type
		maxIndex    int
	)

	if !anyItems() {
		NoItems(w, r)
		return
	}

	id = selectID(params["id"])

	if id == 0 {
		IDNotFound(w, r)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&item)
	handleError(err)

	updateQuery = "UPDATE items SET "

	preInterfacedValues := reflect.ValueOf(item)
	values = make([]interface{}, preInterfacedValues.NumField())
	valueTypes = preInterfacedValues.Type()

	for i := 0; i < preInterfacedValues.NumField(); i++ {
		values[i] = preInterfacedValues.Field(i).Interface()
	}

	maxIndex = len(values) - 1

	for index, value := range values {

		fieldName = strcase.ToSnake(valueTypes.Field(index).Name)

		if fieldName == "id" {
			continue
		}

		updateQuery = updateQuery + fieldName + " = "

		switch value.(type) {
		case int:
			updateQuery = updateQuery + strconv.Itoa(value.(int))
			break
		case string:
			updateQuery = updateQuery + "'" + value.(string) + "'"
			break
		case bool:
			updateQuery = updateQuery + strconv.FormatBool(value.(bool))
			break
		default:
			log.Fatal("Un categorized type found in JSON")
		}

		if index != maxIndex {
			updateQuery = updateQuery + ", "
		}

	}

	updateQuery = updateQuery + " WHERE id = " + params["id"] + ";"

	_, err = db.Query(updateQuery)
	handleError(err)

	createResponse(w, "item_updated")

}

func deleteItemRecord(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var id = selectID(params["id"])

	if id == 0 {
		IDNotFound(w, r)
		return
	}

	_, err := db.Query("DELETE FROM items WHERE id = " + params["id"])
	handleError(err)

	createResponse(w, "item_deleted")

}
