package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloWorld()).Methods("GET")

	log.Fatal(http.ListenAndServe(":8082", router))
}

func helloWorld() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	}
}