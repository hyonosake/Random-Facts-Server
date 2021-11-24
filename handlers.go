package main
//	testing package
//	mock server
// sql injection

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	//"io/ioutil"
	//"encoding/json"
	"log"
)

func	validateId(w http.ResponseWriter, r *http.Request) (id int, err error)	{
	id, err = strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/fact/"))	// better to do it with regexpr
	if err != nil	{
		return id, err
	}
	return id, nil
}

// URL/
func	generalHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)		//<-- http://localhost:8080/abracadabra
		return
	}
	fmt.Fprintf(w, "General Handler\n") // <-- works just fine
	w.WriteHeader(http.StatusAccepted)
}

// URL/fact
func	getHandler(w http.ResponseWriter, r *http.Request)	{
	//var records Records
	switch r.Method	{
	case "GET":
		fmt.Fprintf(w, "Get method called\n")
		//	TODO: Combine all rows in database
		// GET all values into JSON from DB
	case "POST":
		// POST values into DB from JSON
		fmt.Fprintf(w, "POST method called\n")
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusAccepted)
}

//	URL/fact/
func	idHandler(w http.ResponseWriter, r *http.Request)	{
	id, err := validateId(w, r); if err != nil	{
		fmt.Fprintf(w, "Invalid Index\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Get ID=%d method called\n", id)
		h.getUniqueFact(id)
	case "PUT":
		fmt.Fprintf(w, "PUT ID=%d method called\n", id)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	//	TODO: go for Query and, if found index in database, return listing
	w.WriteHeader(http.StatusFound)
}

func (h *RequestHandler) runHandlers()	{

	// Let gv{1,2,3} handle routes{1,2,3} respectively
	//h.sm.HandleFunc("/fact", getHandler)
	//h.sm.HandleFunc("/fact/", getUniqueHandler)		// <-- regex for id (num)
	//h.sm.HandleFunc("/", generalHandler)
	// TODO: default handler for 404
	if err := http.ListenAndServe(":8080", h.sm); err != nil {
		log.Fatal(err)
	}

}

//func initRequestHandling()	{
//
//	h.runHandlers()
//}
