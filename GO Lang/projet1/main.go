package main

import (
	"net/http"
	dictionary "projet1/directory/directories"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	myDictionary := dictionary.NewDictionary()

	router.Use(dictionary.LoggerMiddleware)
	router.Use(dictionary.AuthMiddleware)

	router.HandleFunc("/add", AddHandler(myDictionary)).Methods("POST")
	router.HandleFunc("/get/{word}", GetHandler(myDictionary)).Methods("GET")
	router.HandleFunc("/remove/{word}", RemoveHandler(myDictionary)).Methods("DELETE")
	router.HandleFunc("/list", ListHandler(myDictionary)).Methods("GET")

	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}
