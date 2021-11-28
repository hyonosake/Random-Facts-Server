package main

import (
	"fmt"
	"net/http"
)

type ErrMsg struct {
	Code int		`json:"status"`
	Message	string	`json:"message"`
}

// Respond sends JSONified data and writes HTTP Header
func Respond(w http.ResponseWriter, status int, data interface{}) {

	s.logger.Printf("answered with %v\n", data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		if encodeBody(w, data) != nil	{
			fmt.Println("failed to encode")
		}
	}
}

// RespondErr gives error json message
func RespondErr(w http.ResponseWriter, status int, data string) {

	resp := ErrMsg { Code: status, Message: data }
	Respond(w, status, resp)
}
