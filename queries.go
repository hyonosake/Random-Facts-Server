package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func jsonifyRecord(r Records)  {
	jsonStr, _ := json.Marshal(r)
	fmt.Println(string(jsonStr))

}

func (h *RequestHandler) parseNewFacts(w http.ResponseWriter, r *http.Request)	{
	var records Records
	fmt.Println(records)
}


func (h *RequestHandler) getUniqueFact(id int) (r Records, err error)	{
	rows, err := h.conn.Query(context.Background(),
		"SELECT id, title, description FROM facts WHERE id = $1", id)
	if err != nil	{
		return r, err
	}
	defer rows.Close()
//	for rows.Next() {
//		err := rows.Scan(&r.facts.Id, &r.Title, &r.Description, &r.Links)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//	jsonifyRecord(r)
	return r, nil
}

//func (h *RequestHandler) getAllFacts() (r Record, err error)	{
//
//}
