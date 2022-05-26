package server

import (
	"context"
	"encoding/json"
	"github.com/hyonosake/Random-Facts-Server/internal/types"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

// addRows inserts data into DB
func (s *Server) addRows(ctx context.Context, f *types.PostFactsRequest) (*types.PostFactsResponse, error) {
	batch := &pgx.Batch{}
	query := `INSERT INTO facts (title, description, links) VALUES($1, $2, $3) RETURNING id`

	err := validatePostInfo(f)
	if err != nil {
		return nil, err
	}
	for _, fact := range f.Facts {
		batch.Queue(query, fact.Title, fact.Description, fact.Links)
	}
	br := s.db.SendBatch(ctx, batch)
	br.Close()
	return &types.PostFactsResponse{Ids: len(f.Facts)}, nil
}

// postNewFacts validates and add new facts to DB
func (s *Server) postNewFacts(r *http.Request) (*types.PostFactsResponse, error) {

	var resp *types.PostFactsResponse
	var facts types.PostFactsRequest

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &facts)
	if err != nil {
		return nil, err
	}
	resp, err = s.addRows(context.Background(), &facts)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// getUniqueFact returns specific fact from DB based on ID passed
func (s *Server) getUniqueFact(id int) (resp types.GetFactResponse, err error) {
	query := `SELECT id, title, description, links FROM facts WHERE id = $1`

	row := s.db.QueryRow(context.Background(), query, id)
	err = row.Scan(&resp.Id, &resp.Title, &resp.Description, &resp.Links)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// getRandomFact returns fact corresponding to randomly chosen ID
func (s *Server) getRandomFact(ctx context.Context) (*types.GetFactResponse, error) {
	var resp types.GetFactResponse
	query := `SELECT * FROM facts ORDER BY RANDOM() LIMIT 1;`

	err := s.db.QueryRow(ctx, query).Scan(&resp.Id, &resp.Title, &resp.Description, &resp.Links)
	if err != nil && err.Error() != types.ErrNoRows {
		return nil, err
	}
	return &resp, nil
}

// getAllData pulls all data from DB
func (s *Server) getAllData(ctx context.Context) (*types.GetFactsResponse, error) {
	var facts []types.Fact
	query := `SELECT * FROM facts`

	rows, err := s.db.Query(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var temp = types.Fact{}
		err = rows.Scan(&temp.Id, &temp.Title, &temp.Description, &temp.Links)
		if err != nil {
			s.logger.Error("Unable to unmarshall data from query row", zap.Error(err))
		}
		facts = append(facts, temp)
	}
	return &types.GetFactsResponse{Facts: facts}, nil
}

// putUniqueFact changes specific fact from DB based on ID passed
func (s *Server) putUniqueFact(ctx context.Context, r *http.Request) (*types.PostFactsResponse, error) {
	var fact types.FactRequest
	query := `UPDATE facts SET(title, description, links) = ($1, $2, $3) WHERE id=$4`

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &fact)
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec(ctx, query, fact.Title, fact.Description, fact.Links, fact.Id)
	if err != nil {
		return nil, err
	}
	return &types.PostFactsResponse{}, nil
}
