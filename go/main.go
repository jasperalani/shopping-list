package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
	"time"
)

var DB *sqlx.DB

func main() {

	DB, err = InitDB()
	HandleError(err)

	r := mux.NewRouter()
	r.Use(headerMiddleware, requestMiddleware)

	// Handle all preflight request
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Access-Control-Request-Headers, Access-Control-Request-Method")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	// Endpoints
	var (
		endpoint        = "/"
		endpointInteger = "/{id:[0-9]+}"
	)

	// Route Handlers
	r.HandleFunc(endpoint, CreateItemRecord).Methods("POST")
	r.HandleFunc(endpoint, ReadItemRecord).Methods("GET")
	r.HandleFunc(endpointInteger, ReadItemRecord).Methods("GET")
	r.HandleFunc(endpointInteger, UpdateItemRecord).Methods("PUT")
	r.HandleFunc(endpointInteger, DeleteItemRecord).Methods("DELETE")

	r.NotFoundHandler = http.HandlerFunc(HTTPNotFound)

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":10000", r))

}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func requestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Incoming request ", r.Method)
		next.ServeHTTP(w, r)
	})
}

func InitDB() (*sqlx.DB, error){

	var tries int = 5
	for i := 1; i < (tries + 1); i++ {
		db, err := sqlx.Connect("mysql", "root:password@tcp(127.0.0.1:3306)/shopping-list")
		if err == nil {
			log.Print("Successfully connected to database.")
			return db, nil
		}
		if i == 1 {
			log.Print("Trying to connect to database...")
		}
		log.Print("Attempt " + strconv.Itoa(i))
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Fatal error: could not connect to database")
	return nil, err
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
