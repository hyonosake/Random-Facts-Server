package main

//	testing package
//	mock server
// sql injection

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"io/ioutil"
	//"encoding/json"
	"log"
)

func sendResponse(w http.ResponseWriter, response interface{})  {
	fmt.Printf("Sending JSON Data back for response")
	fmt.Println(response)
	//TODO: http.StatusAccepted
	w.Header().Set("Content-Type", "facts/json")
	json.NewEncoder(w).Encode(response)
}

const secretURL = "/data/another_one/all_of_them/please"


// Get main page with some data (possibly); Acces secretURL
func generalHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var jsonResponse interface{}
	if r.URL.Path == secretURL	{
		jsonResponse, err = h.getAllData(); if err != nil	{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		http.NotFound(w, r)
		return
	}
	sendResponse(w, jsonResponse)
}

// Handle GET and POST for URL/fact
func getHandler(w http.ResponseWriter, r *http.Request) {

	var jsonResponse interface{}
	var err error
	switch r.Method {
	case "GET":
		jsonResponse = h.getRandomFact()
	case "POST":
		jsonResponse, err = h.postNewFacts(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil	{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sendResponse(w, jsonResponse)
}

//	Handle GET and PUT for URL/fact/$id
func idSpecifiedHandler(w http.ResponseWriter, r *http.Request) {

	var jsonResponse interface{}
	id, err := validateId(w, r); if err != nil {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		jsonResponse = h.getUniqueFact(id)
	case "PUT":
		fmt.Println("here")
		jsonResponse, err = h.putUniqueFact(r, id)
	default:
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if err != nil	{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sendResponse(w, jsonResponse)
}


func (h *RequestHandler) runHandlers() {

	if err := http.ListenAndServe(":8080", h.sm); err != nil {
		log.Println("Fatal error encountered: ")
		log.Fatal(err)
	}

}
