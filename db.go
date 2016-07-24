package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var currentID int

var db *sql.DB

// inits db conn
func init() {
	db, err := sql.Open("postgres", "user=recycling dbname=recycling")
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", 4)
	fmt.Println(rows)
	if err != nil {
		log.Fatal(err)
	}
}

//DbFindLocation blah
func DbFindLocation(id int) Location {
	//TODO get id
	// return empty Location if not found
	return Location{}
}

//DbCreateLocation blah
func DbCreateLocation(l Location) Location {
	//TODO create location
	return l
}

//DbDestroyLocation blah
func DbDestroyLocation(id int) error {
	//TODO delete location
	return fmt.Errorf("Could not find Location with id of %d to delete", id)
}

//DbIndexLocation blah
func DbIndexLocation() []Location {
	//TODO return all locations
	return nil
}
