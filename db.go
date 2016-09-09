package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db = gorm.DB{}

// inits db conn
func init() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost user=gorm dbname=gorm sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Location{})
	db.AutoMigrate(&Report{})
	//defer db.Close()

	// Migrate the schemas
}

//DbFindLocation blah
func DbFindLocation(id int) Location {
	location := Location{}
	db.First(&location, id)
	return location
}

//DbCreateLocation blah
func DbCreateLocation(l *Location) *Location {
	db.Create(&l)
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
	db.Find(&locations)
	return locations
}
