package main

import (
	"database/sql"
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/iancoleman/strcase"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

func CreateItemRecord(w http.ResponseWriter, request *http.Request) {

	var (
		item         ItemJSON
		id           int
		name         string
		quantity     int
		generatedSql string
		err          error
	)

	err = json.NewDecoder(request.Body).Decode(&item)
	HandleError(err)

	selectItem := sq.Select("id", "name", "quantity").From("items")
	whereExisting := selectItem.Where(sq.Like{"name": item.Name})
	notCompleted := whereExisting.Where(sq.Eq{"completed": 0})

	generatedSql, _, err = notCompleted.ToSql()
	HandleError(err)

	results := DB.QueryRow(generatedSql)

	err = results.Scan(&id, &name, &quantity)
	HandleError(err)

	if item.Name == name {

		finalQuantity := item.Quantity + quantity

		updateQuantity := sq.Update("items").Set("quantity", strconv.Itoa(finalQuantity))
		whereIDMatches := updateQuantity.Where(sq.Eq{"id": strconv.Itoa(id)})
		generatedSql, _, err = whereIDMatches.ToSql()
		HandleError(err)

		_, err = DB.Query(generatedSql)
		HandleError(err)

		CreateResponse(w, "quantity_increased")

	} else {

		insertInto := sq.Insert("items").Columns("name", "url", "image_url", "person", "quantity")
		values := insertInto.Values(item.Name, item.URL, item.ImageURL, item.Person, strconv.Itoa(item.Quantity))
		generatedSql, _, err = values.ToSql()
		HandleError(err)

		_, err = DB.Query(generatedSql)
		HandleError(err)

		var (
			maxID *sql.Rows
			ID    int
		)

		selectMaxID := sq.Select("MAX(id)").From("items")
		generatedSql, _, err = selectMaxID.ToSql()
		HandleError(err)

		maxID, err = DB.Query(generatedSql) // this query is not safe
		HandleError(err)

		if maxID.Next() {
			err = maxID.Scan(&ID)
			HandleError(err)
		}

		CreateResponse(w, "item_created")

	}

}

func ReadItemRecord(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var (
		item   Item
		item_  ItemJSON
		items  []Item
		items_ []ItemJSON
		query  string
		scope  = len(params) > 0
		err    error
	)

	query = StringEvaluator(scope,
		"SELECT id, name, url, image_url, person, quantity, deleted, completed FROM items WHERE id = "+params["id"]+";",
		"SELECT * FROM items WHERE deleted = 0 AND completed = 0;",
	)

	if scope {

		err = DB.Get(&item, query)
		HandleError(err)

		if item.ID == 0 {
			IDNotFound(w, r)
			return
		}

		item_ = ItemJSON{
			ID:        item.ID,
			Name:      item.Name,
			URL:       item.URL,
			ImageURL:  item.ImageURL,
			Person:    item.Person,
			Quantity:  item.Quantity,
			Deleted:   item.Deleted,
			Completed: item.Completed,
		}

		err = json.NewEncoder(w).Encode(item_)
		HandleError(err)

	} else {

		if !AnyItems() {
			NoItems(w, r)
			return
		}

		err = DB.Select(&items, query)
		HandleError(err)

		//log.Println(items)

		for _, item := range items {
			items_ = append(items_, ItemJSON{
				ID:        item.ID,
				Name:      item.Name,
				URL:       item.URL,
				ImageURL:  item.ImageURL,
				Person:    item.Person,
				Quantity:  item.Quantity,
				Created:   item.Created,
				Deleted:   item.Deleted,
				Completed: item.Completed,
			})
		}

		err = json.NewEncoder(w).Encode(items_)
		HandleError(err)
	}

}

func UpdateItemRecord(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var (
		id           float32
		item         ItemJSON
		fieldName    string
		values       []interface{}
		valueTypes   reflect.Type
		generatedSql string
		err          error
	)

	if !AnyItems() {
		NoItems(w, r)
		return
	}

	id = SelectID(params["id"])

	if id == 0 {
		IDNotFound(w, r)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&item)
	HandleError(err)

	updateItems := sq.Update("items")

	preInterfacedValues := reflect.ValueOf(item)
	values = make([]interface{}, preInterfacedValues.NumField())
	valueTypes = preInterfacedValues.Type()

	for i := 0; i < preInterfacedValues.NumField(); i++ {
		values[i] = preInterfacedValues.Field(i).Interface()
	}

	var setValues sq.UpdateBuilder

	for index, value := range values {

		fieldName = strcase.ToSnake(valueTypes.Field(index).Name)

		if fieldName == "id" {
			continue
		}

		var updateValue string

		switch value.(type) {
		case int:
			updateValue = strconv.Itoa(value.(int))
			break
		case string:
			updateValue = "'" + value.(string) + "'"
			break
		case bool:
			updateValue = strconv.FormatBool(value.(bool))
			break
		default:
			log.Fatal("Unknown type found in JSON")
		}

		if index == 0 {
			setValues = updateItems.Set(fieldName, updateValue)
		} else {
			setValues = setValues.Set(fieldName, updateValue)
		}

	}

	whereIDMatches := setValues.Where(sq.Eq{"id": params["id"]})
	generatedSql, _, err = whereIDMatches.ToSql()
	HandleError(err)

	_, err = DB.Query(generatedSql)
	HandleError(err)

	CreateResponse(w, "item_updated")

}

func DeleteItemRecord(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var id = SelectID(params["id"])

	if id == 0 {
		IDNotFound(w, r)
		return
	}

	_, err := DB.Query("DELETE FROM items WHERE id = " + params["id"])
	HandleError(err)

	CreateResponse(w, "item_deleted")

}
