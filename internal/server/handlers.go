package server

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

//var secretURL = os.Getenv("SECRET_URL")
var secretURL = "/data/another_one/all_of_them/please"

func (s *Server) getJSONResponse(resp interface{}) []byte {
	body, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		s.logger.Error("Unable to marshal data", zap.Error(err))
	}
	return body
}

// HandleRoot handles main page and checks access for secretURL
func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("handleRoot operating fact", zap.String("method", r.Method))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Println("path:", r.URL.Path)
	fmt.Println("path:", secretURL)
	switch r.URL.Path {
	case secretURL:
		resp, err := s.getAllData(ctx)
		if err != nil {
			s.logger.Error("Unable to process data", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(s.getJSONResponse(resp))
		}
	default:
		w.Write([]byte(`Please, specify request path.`))
		w.WriteHeader(http.StatusBadRequest)
	}
}

// HandleFact handles GET and POST for URL/fact
func (s *Server) handleFact(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("handleFact operating fact", zap.Any("method", r.Method))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	switch r.Method {
	case http.MethodGet:
		resp, err := s.getRandomFact(ctx)
		if err != nil {
			s.logger.Error("Unable to process data", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(s.getJSONResponse(resp))
		}
	case http.MethodPost:
		resp, err := s.postNewFacts(r)
		if err != nil {
			s.logger.Error("Unable to process data", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(s.getJSONResponse(resp))
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

//	handleFactId handles GET and PUT for URL/fact/$id
func (s *Server) handleFactId(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("handleFactId operating fact", zap.Any("method", r.Method))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	id, err := validateId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		resp, err := s.getUniqueFact(id)
		if err != nil {
			s.logger.Error("Unable to process data", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(s.getJSONResponse(resp))
		}
	case "PUT":
		resp, err := s.putUniqueFact(ctx, r)
		if err != nil {
			s.logger.Error("Unable to process data", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(s.getJSONResponse(resp))
		}
	default:
		//RespondErr(w, http.StatusBadRequest, "Invalid json message received")
	}
}
