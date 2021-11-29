package main


import (
	"net/http"
)

const secretURL = "/data/another_one/all_of_them/please"


// HandleRoot handles main page and checks access for secretURL
func HandleRoot(w http.ResponseWriter, r *http.Request) {

	s.logger.Printf("recieved %s", r.Method)
	if r.Method != http.MethodGet	{
		RespondErr(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	switch r.URL.Path {
	case secretURL:
		jsonResponse, err := s.getAllData(); if err != nil {
			RespondErr(w, http.StatusInternalServerError, "Unable to read from database")
		} else {
			Respond(w, http.StatusFound, jsonResponse)
		}
	case "/":
		Respond(w, http.StatusOK, "Welcome")
	default:
		RespondErr(w, http.StatusInternalServerError, "Page Not Found")
	}
}


// HandleFact handles GET and POST for URL/fact
func HandleFact(w http.ResponseWriter, r *http.Request) {

	s.logger.Printf("recieved %s", r.Method)
	switch r.Method {
	case http.MethodGet:
		if s.nRows == 0	{
			RespondErr(w, http.StatusBadRequest, "Db is empty")
		}
		response, err := s.getRandomFact(); if err != nil	{
			RespondErr(w, http.StatusBadRequest, "Unable to get random ID")
		} else {
			Respond(w, http.StatusFound, response)
		}
	case http.MethodPost:
		jsonResponse, err := s.postNewFacts(r); if err != nil	{
			RespondErr(w, http.StatusBadRequest, "Unable to POST")
		} else {
			Respond(w, http.StatusCreated, jsonResponse)
		}
	default:
		RespondErr(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}


//	HandleFactId handles GET and PUT for URL/fact/$id
func HandleFactId(w http.ResponseWriter, r *http.Request) 	{

	s.logger.Printf("recieved %s", r.Method)
	var jsonResponse interface{}
	id, err := ValidateId(r); if err != nil {
		RespondErr(w, http.StatusBadRequest, "Request not allowed")
		return
	}
	switch r.Method {
	case "GET":
		jsonResponse, err = s.getUniqueFact(id); if err != nil	{
			RespondErr(w, http.StatusNotFound, "Does not exist")
		} else {
			Respond(w, http.StatusFound, jsonResponse)
		}
	case "PUT":
		jsonResponse, err = s.putUniqueFact(r, id); if err != nil	{
			RespondErr(w, http.StatusBadRequest, "Invalid json message received")
		} else {
			Respond(w, http.StatusFound, jsonResponse)
		}
	default:
		RespondErr(w, http.StatusBadRequest, "Invalid json message received")
	}
}


func (s *server) routes() {
	s.router.HandleFunc("/fact", HandleFact)
	s.router.HandleFunc("/fact/", HandleFactId)		// <-- regex for id (num)
	s.router.HandleFunc("/", HandleRoot)
}

// runHandlers ListenAndServe with ServerMux
func (s *server) runHandlers() {
	s.routes()
	if err := http.ListenAndServe(":80", s.router); err != nil {
		s.logger.Fatalf("Server down: %v\n", err)
	}
}
