package main

import "github.com/jinzhu/gorm"

//Location blah
type Location struct {
	gorm.Model
	ID        uint   `gorm:"primary_key"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Latitude  string `json:"latitutde"`
	Longitude string `json:"longitude"`
}
