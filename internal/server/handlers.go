package server

import (
	"go.uber.org/zap"
	"net/http"
	"os"
)

var secretURL = os.Getenv("SECRET_URL")

// HandleRoot handles main page and checks access for secretURL
func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {

	s.logger.Info("Handler called", zap.String("method", r.Method))
	if r.Method != http.MethodGet {
		RespondErr(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	switch r.URL.Path {
	case secretURL:
		jsonResponse, err := s.getAllData()
		if err != nil {
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
func (s *Server) handleFact(w http.ResponseWriter, r *http.Request) {

	s.logger.Info("received fact", zap.Any("method", r.Method))
	switch r.Method {
	case http.MethodGet:
		response, err := s.getRandomFact()
		_, _ = response, err
		//if err != nil {
		//	RespondErr(w, http.StatusBadRequest, "Unable to get random ID")
		//} else {
		//	Respond(w, http.StatusFound, response)
		//}
	case http.MethodPost:
		jsonResponse, err := s.postNewFacts(r)
		_, _ = jsonResponse, err
		//if err != nil {
		//	RespondErr(w, http.StatusBadRequest, "Unable to POST")
		//} else {
		//	Respond(w, http.StatusCreated, jsonResponse)
		//}
	default:
		RespondErr(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

//	handleFactId handles GET and PUT for URL/fact/$id
func (s *Server) handleFactId(w http.ResponseWriter, r *http.Request) {

	s.logger.Info("Handler called", zap.String("method", r.Method))
	var jsonResponse interface{}
	id, err := ValidateId(r)
	if err != nil {
		RespondErr(w, http.StatusBadRequest, "Request not allowed")
		return
	}
	switch r.Method {
	case "GET":
		jsonResponse, err = s.getUniqueFact(id)
		if err != nil {
			RespondErr(w, http.StatusNotFound, "Does not exist")
		} else {
			Respond(w, http.StatusFound, jsonResponse)
		}
	case "PUT":
		jsonResponse, err = s.putUniqueFact(r)
		if err != nil {
			RespondErr(w, http.StatusBadRequest, "Invalid json message received")
		} else {
			Respond(w, http.StatusFound, jsonResponse)
		}
	default:
		RespondErr(w, http.StatusBadRequest, "Invalid json message received")
	}
}
