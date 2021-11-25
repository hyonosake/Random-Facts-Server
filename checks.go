package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
)

//validateId checks if given ID presents in DB
func	validateId(w http.ResponseWriter, r *http.Request) (id int, err error)	{
	id, err = strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/fact/"))
	if err != nil	{
		return id, err
	}
	return id, nil
}

// ValidatePostInfo validates incoming POST data
func ValidatePostInfo(queries []FactsStructure) error {

	for _, v := range queries	{
		if v.Id != 0 || v.Title == "" || v.Description == ""	{
			return errors.New("Invalid parameters for POST query\n")
		}
	}
	return nil
}

// MaxId Set h.nRows value to amount of rows in DB table
func	(h *RequestHandler) MaxId()	{

	var temp int
	row := h.db.QueryRow(context.Background(), "SELECT CASE WHEN EXISTS (SELECT * FROM facts LIMIT 1) THEN 0" +
		"ELSE 1 END")
	row.Scan(&temp)
	h.isEmpty = temp == 1
	if h.isEmpty	{
		fmt.Println("It's empty tho")
	} else {
		row = h.db.QueryRow(context.Background(), "SELECT max(id) FROM facts")
		row.Scan(&h.nRows)
		fmt.Println("nRows = ", h.nRows)
	}
}

// emptyJson returns interface body for empty response
func emptyJson() interface{}{
	return make(map[string]string)
}