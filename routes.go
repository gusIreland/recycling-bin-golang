package main

import "net/http"

//Route blah
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"LocationIndex",
		"GET",
		"/locations",
		LocationIndex,
	},
	Route{
		"LocationShow",
		"GET",
		"/locations/{locationId}",
		LocationShow,
	},
	Route{
		"LocationCreate",
		"POST",
		"/locations",
		LocationCreate,
	},
	Route{
		"ReportCreate",
		"POST",
		"/locations/{locationId}/reports",
		ReportCreate,
	},
	Route{
		"ReportShow",
		"GET",
		"/reports",
		ReportShow,
	},
}
