package main

//	testing package
//	mock server
// sql injection

import (
	"fmt"
	"net/http"
	//"io/ioutil"
	//"encoding/json"
	"log"
)

const secretURL = "/data/another_one/all_of_them/please"
// URL/
func generalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path)
	if r.URL.Path == secretURL	{
		h.getAllData(w, r)
	} else if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//fmt.Fprintf(w, "General Handler\n") // <-- works just fine
}

// URL/fact
func getHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Get method called\n")
		//	TODO: Combine all rows in database
		h.getRandomFact(w, r)
	case "POST":
		err := h.parseNewFacts(w, r)
		if err != nil {
			fmt.Fprintf(w, "Failed: %v\n", err)
			//w.WriteHeader(http.StatusBadRequest)
		}
	default:
		// TODO: Make valid err value
		fmt.Fprintf(w, "Unknown request %v\n", r.URL)
		w.WriteHeader(http.StatusBadRequest)
	}
	// TODO: http: superfluous response.WriteHeader call from main.getHandler (handlers.go:43)
	//w.WriteHeader(http.StatusAccepted)
}

//	URL/fact/
func idHandler(w http.ResponseWriter, r *http.Request) {
	id, err := validateId(w, r)
	if err != nil {
		//fmt.Fprintf(w, "Invalid Index\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Get ID=%d method called\n", id)
		//h.getUniqueFact(id)
	case "PUT":
		fmt.Fprintf(w, "PUT ID=%d method called\n", id)
	default:
		w.WriteHeader(http.StatusBadRequest)
		// TODO: Make valid err value
		fmt.Fprintf(w, "Unknown request %v\n", r.URL)
		return
	}
	//	TODO: go for Query and, if found index in database, return listing
	w.WriteHeader(http.StatusFound)
}

func (h *RequestHandler) runHandlers() {

	if err := http.ListenAndServe(":8080", h.sm); err != nil {
		log.Println("Fatal error encountered: ")
		log.Fatal(err)
	}

}
