package server

import (
	"net/http"
)

type ErrMsg struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

type Fact struct {
	Id          int      `json:"id,omitempty"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Links       []string `json:"links,omitempty"`
}

type PostFactsResponse struct {
	Ids []int `json:"ids,omitempty"`
}

type PostFactsRequest struct {
	Facts []Fact `json:"facts,omitempty"`
}

type FactRequest = Fact

type GetFactResponse = Fact

// Respond sends JSONified data and writes HTTP Header
func Respond(w http.ResponseWriter, status int, data interface{}) {

	s.logger.Printf("answered with %v\n", data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		if encodeBody(w, data) != nil {
			s.logger.Println("failed to encode")
		}
	}
}

// RespondErr gives error json message
func RespondErr(w http.ResponseWriter, status int, data string) {

	resp := ErrMsg{Code: status, Message: data}
	Respond(w, status, resp)
}
