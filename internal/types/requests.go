package types

type Fact struct {
	Id          int      `json:"id,omitempty"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Links       []string `json:"links,omitempty"`
}

type PostFactsResponse struct {
	Ids int `json:"inserted,omitempty"`
}

type PostFactsRequest struct {
	Facts []Fact `json:"facts,omitempty"`
}

type GetFactsResponse struct {
	Facts []Fact `json:"facts,omitempty"`
}

type FactRequest = Fact

type GetFactResponse = Fact
