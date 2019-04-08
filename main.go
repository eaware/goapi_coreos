package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	router := newRouter()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}
