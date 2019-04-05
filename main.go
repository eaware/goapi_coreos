package main

import (
	"log"
	"net/http"

	"github.com/eaware/goapi_coreos/idk"
)

func main() {
	router := newRouter()

	idk.Idk()

	log.Fatal(http.ListenAndServe(":8000", router))
}
