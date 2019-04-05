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
		"Index",
		"GET",
		"/",
		index,
	},
	route{
		"TodoIndex",
		"GET",
		"/automic",
		todoIndex,
	},
	route{
		"TodoCreate",
		"POST",
		"/automic",
		todoCreate,
	},
	route{
		"TodoShow",
		"GET",
		"/automic/{todoId}",
		todoShow,
	},
}
