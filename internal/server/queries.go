package server

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

// addRows inserts data into DB
func (s *Server) addRows(ctx context.Context, f *PostFactsRequest) (*PostFactsResponse, error) {

	var ids []int
	batch := &pgx.Batch{}
	for _, fact := range f.Facts {
		err := s.db.QueryRow(ctx,
			"INSERT INTO facts (title, description, links) VALUES($1, $2, $3) RETURNING id",
			fact.Title, fact.Description, fact.Links)
		if err != nil {
			s.logger.Error("Unable to insert data", zap.Int("id", fact.Id))
			continue
		}
		ids = append(ids, 5)
	}
	br := s.db.SendBatch(ctx, batch)
	_ = br
	//br.QueryRow().Scan()
	return &PostFactsResponse{Ids: ids}, nil
}

// postNewFacts validates and add new facts to DB
func (s *Server) postNewFacts(r *http.Request) (*PostFactsResponse, error) {

	var resp *PostFactsResponse
	var facts *PostFactsRequest

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, facts)
	if err != nil {
		return nil, err
	}
	resp, err = s.addRows(nil, facts)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// getUniqueFact returns specific fact from DB based on ID passed
func (s *Server) getUniqueFact(id int) (resp GetFactResponse, err error) {
	query := `SELECT id, title, description, links FROM facts WHERE id = $1`

	row := s.db.QueryRow(context.Background(), query, id)
	row.Scan(&resp.Id, &resp.Title, &resp.Description, &resp.Links)
	if resp.Id == 0 {
		return resp, errors.New("Invalid id")
	}
	return resp, nil
}

// getRandomFact returns fact corresponding to randomly chosen ID
func (s *Server) getRandomFact() (*GetFactResponse, error) {
	var resp *GetFactResponse
	query := `SELECT column FROM table ORDER BY RANDOM() LIMIT 1`

	row, err := s.db.Query(context.Background(), query)
	if err != nil {
		//s.logger.Error("Unable to get random data from DB", zap.Error(err))
		return nil, err
	}

	err = row.Scan(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// getAllData pulls all data from DB
func (s *Server) getAllData() ([]GetFactResponse, error) {
	var resp []GetFactResponse
	query := `SELECT * FROM facts`

	rows, err := s.db.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var temp = Fact{}
		err = rows.Scan(&temp.Id, &temp.Title, &temp.Description, &temp.Links)
		if err != nil {
			s.logger.Error("Unable to unmarshall data from query row", zap.Error(err))
		}
		resp = append(resp, temp)
	}
	return resp, nil
}

// putUniqueFact changes specific fact from DB based on ID passed
func (s *Server) putUniqueFact(r *http.Request) (*PostFactsResponse, error) {
	var fact FactRequest
	query := `UPDATE facts SET(title, description, links) = ($1, $2, $3) WHERE id=$4`

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &fact)
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec(context.Background(), query, fact.Title, fact.Description, fact.Links, fact.Id)
	if err != nil {
		return nil, err
	}
	return &PostFactsResponse{}, nil
}
