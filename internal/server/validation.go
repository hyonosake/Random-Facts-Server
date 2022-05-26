package server

import (
	"github.com/hyonosake/Random-Facts-Server/internal/types"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
)

func badQueryData(fact types.FactRequest) bool {
	return fact.Id == 0 || fact.Title == "" || fact.Description == ""
}

func validateId(r *http.Request) (id int, err error) {
	id, err = strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/fact/"))
	if err != nil {
		return id, err
	}
	return id, nil
}

func validatePostInfo(req *types.PostFactsRequest) error {

	for _, fact := range req.Facts {
		if badQueryData(fact) {
			return errors.New("Invalid parameters for POST query\n")
		}
	}
	return nil
}
