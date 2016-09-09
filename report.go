package main

import "github.com/jinzhu/gorm"

//Report blah
type Report struct {
	gorm.Model
	LocationID string `json:"locationId"`
	ReportText string `json:"report"`
}
