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
	Id		int		`json:"id,omitempty"`
	Title	string		`json:"title"`
	Description	string	`json:"description"`
	Links	[]string	`json:"links,omitempty"`
}

var h *RequestHandler

const DATABASE_URL string = "postgres://postgres:root@localhost:5432/randomfacts"


type RequestHandler struct	{
	conn	*pgx.Conn		// db connection
	sm		*http.ServeMux	// handlers
	nRows	int			// Couldn't find how to check last inserted id
	//logging (?)

}

func newRequestHandler(_conn *pgx.Conn, _sm *http.ServeMux)	*RequestHandler  {
	return &RequestHandler	{
		conn: _conn,
		sm: _sm,
	}
}

func initHandling()	{
	connection, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	var sm = http.NewServeMux()
	sm.HandleFunc("/fact", getHandler)
	sm.HandleFunc("/fact/", idHandler)		// <-- regex for id (num)
	sm.HandleFunc("/", generalHandler)
	h = newRequestHandler(connection, sm)
}

func testJson()	{
}

func main()	{
	initHandling()
	defer h.conn.Close(context.Background())
	log.Printf("Service started\n")
	testJson()
	h.runHandlers()
	//initRequestHandling()

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