package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
)


//ValidateId checks if given ID presents in DB
func	ValidateId(r *http.Request) (id int, err error)	{
	id, err = strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/fact/"))
	if err != nil	{
		return id, err
	}
	return id, nil
}

// validatePostInfo validates incoming POST data
func validatePostInfo(queries []FactsStructure) error {

	for _, v := range queries	{
		if v.Id != 0 || v.Title == "" || v.Description == ""	{
			return errors.New("Invalid parameters for POST query\n")
		}
	}
	return nil
}

// MaxId Set h.nRows value to amount of rows in DB table
func (s *server) MaxId()	{

	var temp int
	row := s.db.QueryRow(context.Background(), "SELECT CASE WHEN EXISTS (SELECT * FROM facts LIMIT 1) THEN 0" +
		"ELSE 1 END")
	row.Scan(&temp)
	s.isEmpty = temp == 1
	if s.isEmpty != false	{
		row = s.db.QueryRow(context.Background(), "SELECT max(id) FROM facts")
		row.Scan(&s.nRows)
	}
}

// EmptyJson returns interface body for empty response
func EmptyJson() interface{}{
	return make(map[string]string)
}


func DecodeBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func encodeBody(w http.ResponseWriter, data  interface{}) error {
	fmt.Printf("encoding %v\n",  data)
	return json.NewEncoder(w).Encode(data)
}

// mapForIds creates map {key : values}
func mapForIds(key string, values []int) (k map[string][]int) {

	k = make(map[string][]int)
	k[key] = values
	return
}

func badQueryData(fact *FactsStructure, urlid int)	bool	{
	return fact.Id == 0 || fact.Id != urlid || fact.Description == "" || fact.Title == ""

}