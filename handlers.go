package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//Index blah
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You should recycle more")
}

//LocationIndex blah
func LocationIndex(w http.ResponseWriter, r *http.Request) {
	locations := DbIndexLocation()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(locations); err != nil {
		panic(err)
	}
}

//LocationShow blah
func LocationShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	locationID := vars["locationId"]
	fmt.Fprintln(w, "Location show:", locationID)
}

//LocationCreate blah
func LocationCreate(w http.ResponseWriter, r *http.Request) {
	var location Location
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &location); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err2 := json.NewEncoder(w).Encode(err); err2 != nil {
			panic(err)
		}
	}

	l := DbCreateLocation(location)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(l); err != nil {
		panic(err)
	}
}

//ReportCreate blah
func ReportCreate(w http.ResponseWriter, r *http.Request) {
	var location Location
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &location); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err2 := json.NewEncoder(w).Encode(err); err2 != nil {
			panic(err)
		}
	}

	l := DbCreateLocation(location)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(l); err != nil {
		panic(err)
	}
}

//ReportShow blah
func ReportShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	locationID := vars["locationId"]
	fmt.Fprintln(w, "Location show:", locationID)
}
