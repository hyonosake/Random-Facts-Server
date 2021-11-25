package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
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

var h *RequestHandler

const DbFacts string = "postgres://postgres:root@localhost:5432/randomfacts"
const DbLog string = "postgres://postgres:root@localhost:5432/randomfacts"


type RequestHandler struct	{
	db		*pgx.Conn		// db connection
	logs	*pgx.Conn		// logs data
	sm		*http.ServeMux	// handlers
	nRows	int				// index of last inserted row
	isEmpty	bool			// check if anything in table

	//logging (?)
}

func newRequestHandler(_conn *pgx.Conn, _sm *http.ServeMux, _ *pgx.Conn)	*RequestHandler  {
	return &RequestHandler	{
		//logs: logdb,
		db: _conn,
		sm: _sm,
	}
}

func initHandling()	{
	connection, err := pgx.Connect(context.Background(), DbFacts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	logs, err := pgx.Connect(context.Background(), DbLog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	var sm = http.NewServeMux()
	sm.HandleFunc("/fact", getHandler)
	sm.HandleFunc("/fact/", idSpecifiedHandler)		// <-- regex for id (num)
	sm.HandleFunc("/", generalHandler)
	h = newRequestHandler(connection, sm, logs)
}

func main()	{
	initHandling()
	defer h.db.Close(context.Background())
	//defer h.logs.Close(context.Background())
	log.Printf("Service started\n")
	h.MaxId()
	fmt.Println(h.nRows)
	//var any interface{}
	//any = 5
	//any = 2.4
	//fmt.Println(any)
	h.runHandlers()
	// to close DB pool


	//row, err := conn.Query(context.Background(), "SELECT id, title, description FROM facts WHERE id = $1", 2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for row.Next() {
	//	// TODO: append to array of recs and Marshall them
	//	// Scan reads the values from the current row into rec
	//	row.Scan(&rec.Id, &rec.Title, &rec.Description)
	//	fmt.Printf("%+v\n", rec)
	//	// Records.append()
	//	//
	//	// Marshal into JSON
	//}
	//defer row.Close()


	// usually, this is taken as an environment variable as in below commented out code
	// databaseUrl = os.Getenv("DATABASE_URL")
	// for the time being, let's hard code it as follows.
	// ensure to change values as needed.
	// this returns connection pool

	//res, err := conn.Exec (context.Background(),
		//"INSERT INTO facts(title, description) VALUES($1, $2)", "без ссылок", "некоторые факты не имеют" +
	//if err != nil	{
	//	fmt.Fprintf(os.Stderr, "Unable to Exec: %v\n", err)
	//	os.Exit(123)
	//}
}