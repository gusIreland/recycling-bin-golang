package main

import "github.com/jinzhu/gorm"

//Report blah
type Report struct {
	gorm.Model
	ID         string `json:"id"`
	LocationID string `json:"locationId"`
	ReportText string `json:"report"`
}
