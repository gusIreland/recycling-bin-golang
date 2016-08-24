package main

import (
	"log"
	"net/http"
	//"github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
