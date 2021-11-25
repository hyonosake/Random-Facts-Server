package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"net/http"
	"os"
)

type PostQuery map[string][]FactsStructure

type FactsStructure struct	{
	Id		int			`json:"id,omitempty"`
	Title	string		`json:"title"`
	Description	string	`json:"description"`
	Links	[]string	`json:"links,omitempty"`
}

type RequestHandler struct	{
	db		*pgx.Conn		// db connection
	sm		*http.ServeMux	// handlers
	nRows	int				// index of last inserted row
	isEmpty	bool			// check if anything in table


}

var h *RequestHandler

const DbFacts string = "postgres://postgres:root@localhost:5432/randomfacts"
const secretURL = "/data/another_one/all_of_them/please"


// newRequestHandler links handler and data
func newRequestHandler(_conn *pgx.Conn, _sm *http.ServeMux)	*RequestHandler  {
	return &RequestHandler	{
		db: _conn,
		sm: _sm,
	}
}

// initHandling connects to DB and link paths to handler functions
func initHandling()	{
	connection, err := pgx.Connect(context.Background(), DbFacts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	h = newRequestHandler(connection, http.NewServeMux())
	h.sm.HandleFunc("/fact", getHandler)
	h.sm.HandleFunc("/fact/", idSpecifiedHandler)		// <-- regex for id (num)
	h.sm.HandleFunc("/", generalHandler)
	h.MaxId()
}


func main()	{
	initHandling()
	defer h.db.Close(context.Background())
	h.runHandlers()
}