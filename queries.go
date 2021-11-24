package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func jsonifyRecord(r Records)  {
	jsonStr, _ := json.Marshal(r)
	fmt.Println(string(jsonStr))
}


func (h *RequestHandler) parseNewFacts(w http.ResponseWriter, r *http.Request)	{

	body, err := ioutil.ReadAll(r.Body); if err != nil {
		log.Println("Error in body, err")
	}
	var facts map[string]Records
	err = json.Unmarshal([]byte(body), &facts); if err != nil {
		//TODO: check what if returns error
		fmt.Fprintf(w, "Invalid POST request body\n")
	}
	fmt.Println(facts["facts"])
	for i, v := range facts["facts"]	{
		fmt.Println(i, ":", v)
	fmt.Println(facts)
}

// TODO: JSONIFY
func (h *RequestHandler) getUniqueFact(id int) (r int, err error)	{
	rows, err := h.conn.Query(context.Background(),
		"SELECT id, title, description FROM facts WHERE id = $1", id)
	if err != nil	{
		return r, err
	}
	//defer rows.Close()
	//for rows.Next() {
	//	err := rows.Scan(&r.facts.Id, &r.Title, &r.Description, &r.Links)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//jsonifyRecord(r)
	return r, nil
}

//func (h *RequestHandler) getAllFacts() (r Record, err error)	{
//
//}
