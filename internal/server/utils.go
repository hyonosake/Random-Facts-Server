package server

import (
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
)

func badQueryData(fact FactRequest) bool {
	return fact.Id == 0 || fact.Title == "" || fact.Description == ""
}

//ValidateId checks if given ID presents in DB
func ValidateId(r *http.Request) (id int, err error) {
	id, err = strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/fact/"))
	if err != nil {
		return id, err
	}
	return id, nil
}

// validatePostInfo validates incoming POST data
func validatePostInfo(facts []FactRequest) error {

	for _, fact := range facts {
		if badQueryData(fact) {
			return errors.New("Invalid parameters for POST query\n")
		}
	}
	return nil
}
