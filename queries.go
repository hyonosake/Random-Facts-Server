package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"io/ioutil"
	"log"
	"net/http"
)

// Creates map {key : values}
func mapForIds(key string, values []int) (k map[string][]int) {

	k = make(map[string][]int)
	k[key] = values
	return
}

// Function for adding data into DB
func (h *RequestHandler) addToData(queries []FactsStructure) (val []int, err error) {

	var lastId int
	for i := range queries {
		err := h.db.QueryRow(context.Background(),
			"INSERT into facts (title, description, links) VALUES($1, $2, $3) RETURNING id",
			queries[i].Title, queries[i].Description, queries[i].Links).Scan(&lastId)
		if err != nil {
			return val, err
		}
		val = append(val, lastId)
		h.nRows = lastId
		log.Printf("INSERT id=%d;title=\"%s\";description=\"%s\";links=%v",
			lastId, queries[i].Title, queries[i].Description, queries[i].Links)
	}
	return val, nil
}


// Validate and add new facts to DB
func (h *RequestHandler) postNewFacts(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	var values []int
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return values, err
	}
	var facts PostQuery
	err = json.Unmarshal([]byte(body), &facts)
	if err != nil {
		fmt.Println("Hey im here btw")
		return values, err
	}
	err = ValidatePostInfo(facts["facts"])
	if err != nil {
		return values, err
	}
	values, err = h.addToData(facts["facts"])
	if err != nil {
		return values, err
	}
	return mapForIds("ids", values), nil
}


//TODO: What if there are 0 rows ?
func (h *RequestHandler) getRandomFact() interface{}	{

	response := FactsStructure{}
	if h.isEmpty == false {
		rand.Seed(time.Now().UnixNano())
		id := rand.Intn(h.nRows)
		if id == 0	{
			id += 1
		}
		fmt.Println("Chose ", id)
		row := h.db.QueryRow(context.Background(), "SELECT * FROM facts WHERE id=$1", id)
		row.Scan(&response.Id, &response.Title, &response.Description, &response.Links)
		fmt.Printf("%v\n", response)
	}
	return response
}


//TODO: Error return maybe ?
func (h *RequestHandler) getAllData() (any interface{}, err error) {

	result := make([]FactsStructure, 1)	// TODO: check
	rows, err := h.db.Query(context.Background(), "SELECT * FROM facts")
	if err != nil {
		return make(map[string]string), err
	}
	defer rows.Close()
	for rows.Next() {
		var response FactsStructure
		err := rows.Scan(&response.Id, &response.Title, &response.Description, &response.Links)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, response)
	}
	if len(result) == 0	{
		return make(map[string]string), nil
	}
	return result, nil
}


// returns specific fact from DB based on ID passed
func (h *RequestHandler) getUniqueFact(id int) (any FactsStructure)	{

	row := h.db.QueryRow(context.Background(),
		"SELECT id, title, description, links FROM facts WHERE id = $1", id)
	row.Scan(&any.Id, &any.Title, &any.Description, &any.Links)
	return
}


// Changes specific fact from DB based on ID passed
func (h *RequestHandler) putUniqueFact(r *http.Request, id int) (any interface{}, err error)	{

	body, err := ioutil.ReadAll(r.Body)
	any = make(map[string]string)
	if err != nil {
		fmt.Println("1st err")
		return any, err
	}
	var fact FactsStructure
	err = json.Unmarshal([]byte(body), &fact)
	if err != nil{
		fmt.Println("2nd err")
		return any, err
	}
	if fact.Title == "" || fact.Description == ""	|| fact.Id != id {
		fmt.Println("3rd err")
		return any, err
	}
	_, err = h.db.Exec(context.Background(),
		"UPDATE facts SET(title, description, links) = ($1, $2, $3) WHERE id=$4",
		fact.Title, fact.Description, fact.Links, fact.Id)
	if err != nil {
		fmt.Println("4th err")
		return any, err
	}
	fmt.Println("empty err")
	return any, nil
}
