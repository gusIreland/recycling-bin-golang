package main

//Location blah
type Location struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Score     int      `json:"score"`
	Address   string   `json:"address"`
	Latitude  string   `json:"latitutde"`
	Longitude string   `json:"ongitude"`
	Reports   []Report `json:"reports"`
}
