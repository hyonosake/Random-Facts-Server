package main

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)


// addToData inserts data into DB
func (s *server) addToData(queries []FactsStructure) (val []int, err error) {

	var lastId int
	for i := range queries {
		err := s.db.QueryRow(context.Background(),
			"INSERT into facts (title, description, links) VALUES($1, $2, $3) RETURNING id",
			queries[i].Title, queries[i].Description, queries[i].Links).Scan(&lastId)
		if err != nil {
			return val, err
		}
		val = append(val, lastId)
		s.nRows = lastId
	}
	return val, nil
}


// postNewFacts validates and add new facts to DB
func (s *server) postNewFacts(r *http.Request) (interface{}, error) {

	var values []int
	var facts PostQuery

	body, err := ioutil.ReadAll(r.Body); if err != nil {
		return values, err
	}
	defer r.Body.Close()
	err = json.Unmarshal([]byte(body), &facts)
	if err != nil {
		return values, err
	}
	err = validatePostInfo(facts["facts"]); if err != nil {
		return values, err
	}
	values, err = s.addToData(facts["facts"]); if err != nil {
		return values, err
	}
	return mapForIds("ids", values), nil
}


// getUniqueFact returns specific fact from DB based on ID passed
func (s *server) getUniqueFact(id int) (any FactsStructure, err error) {

	row := s.db.QueryRow(context.Background(),
		"SELECT id, title, description, links FROM facts WHERE id = $1", id)
	row.Scan(&any.Id, &any.Title, &any.Description, &any.Links)
	if any.Id == 0	{
		return any, errors.New("Invalid id")
	}
	return any, nil
}


// getRandomFact returns fact corresponding to randomly chosen ID
func (s *server) getRandomFact() (response interface{}, err error)	{

	//var response  FactsStructure
	if s.isEmpty == false {
		rand.Seed(time.Now().UnixNano())
		id := rand.Intn(s.nRows)
		if id == 0	{
			id += 1
		}
		response, err := s.getUniqueFact(id); if err != nil	{
			return response, errors.New("Unable to get chosen id")
		}
	}
	return response, nil
}


// getAllData pulls all data from DB
func (s *server) getAllData() (any interface{}, err error) {

	var result []FactsStructure
	rows, err := s.db.Query(context.Background(), "SELECT * FROM facts"); if err != nil {
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


// putUniqueFact changes specific fact from DB based on ID passed
func (s *server) putUniqueFact(r *http.Request, id int) (any interface{}, err error)	{

	any = make(map[string]string)
	var fact FactsStructure
	body, err := ioutil.ReadAll(r.Body); if err != nil	{
		return any, err
	}
	err = json.Unmarshal([]byte(body), &fact); if err != nil	{
		return any, errors.New("Unable to unmarshal JSON")
	}
	if fact.Title == "" || fact.Description == ""	|| fact.Id != id || fact.Id == 0	{
		return any, errors.New("invalid id")
	}
	_, err = s.db.Exec(context.Background(),
		"UPDATE facts SET(title, description, links) = ($1, $2, $3) WHERE id=$4",
		fact.Title, fact.Description, fact.Links, fact.Id)
	if err != nil {
		return any, errors.New("Unable to connect to DB")
	}
	return any, nil
}
