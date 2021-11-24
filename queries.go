package main

import (
	"context"
	"encoding/json"
	//"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
)

//TODO: return value of a function
func (h *RequestHandler) addToData(queries []FactsStructure) (val []int, err error) {

	var lastId int
	for i := range queries {
		err := h.conn.QueryRow(context.Background(),
			"INSERT into facts (title, description, links) VALUES($1, $2, $3) RETURNING id",
			queries[i].Title, queries[i].Description, queries[i].Links).Scan(&lastId)
		if err != nil {
			return val, err
		}
		val = append(val, lastId)
		log.Printf("INSERT id=%d;title=\"%s\";description=\"%s\";links=%v",
			lastId, queries[i].Title, queries[i].Description, queries[i].Links)
	}
	return val, nil
}

func (h *RequestHandler) parseNewFacts(w http.ResponseWriter, r *http.Request) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.New("Bad request")
	}
	var facts PostQuery
	err = json.Unmarshal([]byte(body), &facts)
	if err != nil {
		//TODO: check what if returns error. Superfluous response.WriteHeader call from here
		return errors.New("Invalid POST request body")
	}
	err = ValidatePostInfo(facts["facts"])
	if err != nil {
		return errors.New("Invalid POST request body")
	}
	val, err := h.addToData(facts["facts"])
	if err != nil {
		return errors.New("Unable to add to data")
	}
	// TODO: What's that?
	w.Header().Set("Content-Type", "facts/json")
	//TODO: return { "ids": val }. Is it ok to do it here?
	json.NewEncoder(w).Encode(jsoinfyPostRequest(val))
	return nil
}

// TODO: JSONIFY

func jsoinfyPostRequest(values []int) (k map[string][]int) {

	k = make(map[string][]int)
	k["ids"] = values
	return
}

func (h *RequestHandler) jsonifyAllData() {

	//response
	//rows, err := h.conn.Query(context.Background(), "SELECT * FROM facts")
	//if err != nil {
	//	return
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	err := rows.Scan(&r.facts.Id, &r.Title, &r.Description, &r.Links)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	////jsonifyRecord(r)
	return
}

//func (h *RequestHandler) getUniqueFact(id int) ()	{
//	rows, err := h.conn.Query(context.Background(),
//		"SELECT id, title, description FROM facts WHERE id = $1", id)
//	if err != nil	{
//		return r, err
//	}
//	defer rows.Close()
//	//for rows.Next() {
//	//	err := rows.Scan(&r.facts.Id, &r.Title, &r.Description, &r.Links)
//	//	if err != nil {
//	//		log.Fatal(err)
//	//	}
//	//}
//	//jsonifyRecord(r)
//	return r, nil
//}
