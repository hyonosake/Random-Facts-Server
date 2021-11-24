package main

import (
	"context"
	//"net/http"
	//"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	// pq
	"log"
	"os"
	"time"
)

type Record struct {
	Id		int32		`json:"id"`
	Title	string		`json:"title"`
	Description	string	`json:"description"`
	Link	[]string	`json:"link"`
}

type Records []Record

const DATABASE_URL string = "postgres://postgres:root@localhost:5432/randomfacts"


func getFromJson() {
	type FruitBasket struct {
		Name    string
		Fruit   []string
		Id      int64 `json:"ref"`
		Created time.Time
	}

	jsonData := []byte(`
{
    "Name": "Standard",
    "Fruit": [
        "Apple",
        "Banana",
        "Orange"
    ],
    "ref": 999,
    "Created": "2018-04-09T23:00:00Z"
}`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(basket.Name, basket.Fruit, basket.Id)
	fmt.Println(basket.Created)
}

type HandleConnection struct	{
	connection *Conn

	// connection
	// logging (?)
	
}

func main()	{

	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	var rec Record
	//fmt.Println(rec)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// to close DB pool
	defer conn.Close(context.Background())


	row, err := conn.Query(context.Background(), "SELECT id, title, description FROM facts WHERE id = $1", 2)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		// TODO: append to array of recs and Marshall them
		// Scan reads the values from the current row into rec
		row.Scan(&rec.Id, &rec.Title, &rec.Description)
		fmt.Printf("%+v\n", rec)
		// Records.append()
		//
		// Marshal into JSON
	}
	//fmt.Println(rec.Id, rec.Title, rec.Description)
	defer row.Close()



	//defer rows.Close()
	initRequestHandling()
	// get the database connection URL.
	// usually, this is taken as an environment variable as in below commented out code
	// databaseUrl = os.Getenv("DATABASE_URL")
	// for the time being, let's hard code it as follows.
	// ensure to change values as needed.
	// this returns connection pool

	fmt.Fprintf(os.Stderr, "Congrats! connected to db\n")
	res, err := conn.Exec (context.Background(),
		"INSERT INTO facts(title, description) VALUES($1, $2)", "без ссылок", "некоторые факты не имеют" +
		"ссылок на дополнительную информацию")
	if err != nil	{
		fmt.Fprintf(os.Stderr, "Unable to Exec: %v\n", err)
		os.Exit(123)
	}
	fmt.Fprintf(os.Stderr,"%v\n", res)
}