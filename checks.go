package main

import (
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