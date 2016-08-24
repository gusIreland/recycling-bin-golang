package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

// inits db conn
func init() {
	db, err := gorm.Open("sqlite3", "recycling.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Migrate the schemas
	db.AutoMigrate(&Location{})
	db.AutoMigrate(&Report{})
}

//DbFindLocation blah
func DbFindLocation(id int) Location {
	location := Location{}
	db.First(&location, id)
	return location
}

//DbCreateLocation blah
func DbCreateLocation(l Location) Location {
	db.Create(l)
	return l
}

//DbDestroyLocation blah
func DbDestroyLocation(id int) error {
	location := Location{}
	db.First(&location, id)
	db.Delete(&location)
	return nil
}

//DbIndexLocation blah
func DbIndexLocation() []Location {
	locations := []Location{}
	db.Select("").Find(&locations)
	return nil
}
