package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"os"
)

type PostQuery map[string][]FactsStructure

type FactsStructure struct	{
	Id			int			`json:"id,omitempty"`
	Title		string		`json:"title"`
	Description	string		`json:"description"`
	Links		[]string	`json:"links,omitempty"`
}

type server struct	{
	db		*pgx.Conn		// db connection
	router	*http.ServeMux	// handlers
	logger	*log.Logger		// logs
	nRows	int				// index of last inserted row
	isEmpty	bool			// check if anything in table
}

var s *server

const DbFacts string = "postgres://postgres:root@database:5432/postgres"
const DbFactsLhost string = "postgres://postgres:root@localhost:5432/postgres"


// newRequestHandler links handler and data
func newRequestHandler(_conn *pgx.Conn, _router *http.ServeMux)	*server  {
	return &server	{
		db: _conn,
		router: _router,
	}
}

// initHandling connects to DB and link paths to handler functions
func initHandling()	{

	connection, err := pgx.Connect(context.Background(), DbFacts)
	s = newRequestHandler(connection, http.NewServeMux())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	s.MaxId()
}


func main()	{
	initHandling()
	f, err := os.OpenFile("API_logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	s.logger = log.New(f, "API: ", log.LstdFlags)
	defer f.Close()
	defer s.db.Close(context.Background())
	s.runHandlers()
}