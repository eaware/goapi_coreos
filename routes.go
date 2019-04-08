package main

import "net/http"

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routesT []route

var routes = routesT{
	route{
		"TodoCreate",
		"POST",
		"/automic",
		todoCreate,
	},
}
