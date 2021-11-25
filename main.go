package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
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
	logs	*logrus.Logger	// logs ( ? )
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
		logs: logrus.New(),
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
	h.MaxId()
	var sm = http.NewServeMux()
	sm.HandleFunc("/fact", getHandler)
	sm.HandleFunc("/fact/", idSpecifiedHandler)		// <-- regex for id (num)
	sm.HandleFunc("/", generalHandler)
	h = newRequestHandler(connection, sm)
}


func main()	{
	initHandling()
	defer h.db.Close(context.Background())
	//logFile, err := os.Getwd(); if err != nil	{
	//	log.Fatal()
	//}
	//logFile += "/logs.log"
	//os.Open(logFile)
	h.runHandlers()
}