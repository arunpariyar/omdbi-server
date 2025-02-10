package models

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type SearchResult struct {
	Search []Movie `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response     string `json:"Response"`
}