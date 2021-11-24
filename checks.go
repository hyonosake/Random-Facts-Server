package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
)

func	validateId(w http.ResponseWriter, r *http.Request) (id int, err error)	{
	id, err = strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/fact/"))
	if err != nil	{
		return id, err
	}
	return id, nil
}

// validate incoming data
func ValidatePostInfo(queries []FactsStructure) error {

	for _, v := range queries	{
		if v.Id != 0 || v.Title == "" || v.Description == ""	{
			return errors.New("Invalid parameters for POST query\n")
		}
	}
	return nil
}

func	(h *RequestHandler) MaxId()	{

	row := h.conn.QueryRow(context.Background(), "SELECT CASE" +
		"WHEN EXISTS (SELECT * FROM facts LIMIT 1) THEN 0" +
		"ELSE 1 END;")
	row.Scan(&h.isEmpty)
	if h.isEmpty	{
		fmt.Println("It's empty tho")
	} else {
		row = h.conn.QueryRow(context.Background(), "SELECT max(id) FROM facts")
		row.Scan(&h.nRows)
		fmt.Println("nRows = ", h.nRows)
	}
}