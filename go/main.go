package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	//Init router
	r := mux.NewRouter()

	//Router Handlers / Endpoints
	r.HandleFunc("/", createItemRecord).Methods("POST")
	r.HandleFunc("/", readItemRecord).Methods("GET")
	r.HandleFunc("/{id}", readItemRecord).Methods("GET")
	r.HandleFunc("/{id}", updateItemRecord).Methods("PUT")
	r.HandleFunc("/{id}", deleteItemRecord).Methods("DELETE")

	r.NotFoundHandler = http.HandlerFunc(HTTPNotFound)

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":10000", r))

}

func createDatabase() {

	db, err := sql.Open("mysql", "root:password@tcp(mariadb:3306)/shoppinglist")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE shoppinglist")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE shoppinglist")
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadFile("../table.sql")

	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	fmt.Println(text)

	//    _,err = db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
	//    if err != nil {
	//        panic(err)
	//    }
}

/*

{
    "name": "Socks",
    "url": "https://bit.ly/2WU5xgf",
    "image_url": "https://bit.ly/2ZsyjGt",
    "person": "Jasper",
    "quantity": 2
}

{
    "name": "Socks",
    "url": "https%3A%2F%2Fbit.ly%2F2WU5xgf",
    "image_url": "https%3A%2F%2Fbit.ly%2F2ZsyjGt",
    "person": "Jasper",
    "quantity": 2os
}


{
    "name": "Socks",
    "url": "url",
    "image_url": "image_url",
    "person": "Jasper",
    "quantity": 2
}




*/
