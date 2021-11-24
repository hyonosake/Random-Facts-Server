package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//TODO: return value of a function
func (h *RequestHandler) addToData(queries []FactsStructure) (val []int, err error)  {

	fmt.Printf("AHAHAHHA IM HERE!\n")
	// TODO: get last id that exists in table
	for i := range queries	{
		//_, err = h.conn.Exec(context.Background(), "INSERT into facts (title, description, links)" +
		//	"VALUES($1, $2, $3) RETURNING id", queries[i].Title, queries[i].Description, queries[i].Links)
		//if err != nil	{
		//	//TODO failed to execute
		//	return val, err
		//}
		//h.nRows++
		//val = append(val, i + h.nRows)
		//log.Printf("Inserted at index %d\n", h.nRows)
		log.Printf("INSERT id=%d;title=\"%s\";description=\"%s\";links=%v",
			h.nRows, queries[i].Title, queries[i].Description, queries[i].Links)
	}
	return val, nil
}

func (h *RequestHandler) parseNewFacts(w http.ResponseWriter, r *http.Request)	{

	body, err := ioutil.ReadAll(r.Body); if err != nil {
		log.Printf("Error in body: %s\n", err)
	}
	//var facts map[string][]FactsStructure
	var facts PostQuery
	err = json.Unmarshal([]byte(body), &facts); if err != nil {
		//TODO: check what if returns error
		fmt.Fprintf(w, "Invalid POST request body\n")
	}
	err = ValidatePostInfo(facts["facts"]); if err != nil	{
		log.Printf("Error in post validation: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
		//TODO check this stuff
	}
	val, err := h.addToData(facts["facts"]); if err != nil	{
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(val)
}

// TODO: JSONIFY
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