package main

//	testing package
//	mock server
// sql injection

import (
	"encoding/json"
	"fmt"
	//"github.com/sirupsen/logrus" ( ? )
	"log"
	"net/http"
)

type ResponseErr struct	{
	code int		`json:"status"`
	message string	`json:"message"`
}

func sendResponseErr(w http.ResponseWriter, whatError int, msg string)	{

	fmt.Printf("Sending JSON ERR RESPONSE: ")
	fmt.Fprintln(w, msg)
	w.WriteHeader(whatError)
}

// sendResponse sends response with JSONified response interface or error otherwise
func sendResponse(w http.ResponseWriter, response interface{}, whatHeader int)  {

	fmt.Printf("Sending JSON Data back for response")
	fmt.Println(response)
	//w.WriteHeader(whatHeader)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response); if err != nil {
		sendResponseErr(w, http.StatusInternalServerError, "unable to encode data as json")
	}

}


// generalHandler handles main page and checks access for secretURL
func generalHandler(w http.ResponseWriter, r *http.Request) {

	var err error
	var jsonResponse interface{}
	switch r.URL.Path {
	case secretURL:
		jsonResponse, err = h.getAllData();
		if err != nil {
			sendResponseErr(w, http.StatusInternalServerError, "Unable to read from database")
		}
		sendResponse(w, jsonResponse, http.StatusFound)
	case "/":
		sendResponse(w, jsonResponse, http.StatusOK)
	default:
		fmt.Println("Im here now")
		sendResponseErr(w, http.StatusNotFound, "Page Not Found")
	}
}


// getHandler handles GET and POST for URL/fact
func getHandler(w http.ResponseWriter, r *http.Request) {

	var jsonResponse interface{}
	var err error
	var whatResponse = http.StatusAccepted
	switch r.Method {
	case "GET":
		whatResponse = http.StatusOK
		jsonResponse = h.getRandomFact()
	case "POST":
		whatResponse = http.StatusCreated
		jsonResponse, err = h.postNewFacts(w, r); if err != nil	{
			sendResponseErr(w, http.StatusBadRequest, "Invalid json message received")
			return
		}
	default:
		sendResponseErr(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	sendResponse(w, jsonResponse, whatResponse)
}


//	idSpecifiedHandler handles GET and PUT for URL/fact/$id
func idSpecifiedHandler(w http.ResponseWriter, r *http.Request) {

	var jsonResponse interface{}
	id, err := validateId(w, r); if err != nil {
		sendResponseErr(w, http.StatusBadRequest, "Invalid json message received")
		return
	}
	switch r.Method {
	case "GET":
		jsonResponse = h.getUniqueFact(id)
	case "PUT":
		jsonResponse, err = h.putUniqueFact(r, id); if err != nil	{
			sendResponseErr(w, http.StatusNotModified, "Invalid put parameters")
		}
	default:
		sendResponseErr(w, http.StatusBadRequest, "Invalid json message received")
		return
	}
	sendResponse(w, jsonResponse, http.StatusFound)
}


// runHandlers ListenAndServe with ServerMux
func (h *RequestHandler) runHandlers() {

	if err := http.ListenAndServe(":8080", h.sm); err != nil {
		log.Println("Fatal error encountered: ")
		log.Fatal(err)
	}
}
