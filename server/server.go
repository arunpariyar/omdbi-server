package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arunpariyar/omdbi-server/models"
	"github.com/arunpariyar/omdbi-server/utils"
)


type Server struct {
	cache map[string]models.SearchResult
	apiKey string
	baseUrl string
}


func NewServer(apiKey string) *Server {
	return &Server{
		cache: make(map[string]models.SearchResult),
		apiKey: apiKey,	
		baseUrl: "http://www.omdbapi.com/?apikey=",
	}
}

// searchQuery function to search for a movie by title
func (s *Server)SearchQuery(query string) (models.SearchResult, error) {

	if value, exists := s.cache[query]; exists {
		log.Println("served from cache")
		return value, nil
	}

	if s.apiKey == "" {
		return models.SearchResult{}, fmt.Errorf("API key not found")
	}

	queryURL := s.baseUrl + s.apiKey + "&s=" + query

	response, err := http.Get(queryURL)
	if err != nil {
		return models.SearchResult{}, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.SearchResult{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	searchResult, err := utils.JsonDecoder(response)
	if err != nil {
		return models.SearchResult{}, fmt.Errorf("failed to decode json: %w", err)
	}

	s.cache[query] = searchResult
	return searchResult, nil
}

// searchByTitle function to search for a movie by title
func (s *Server)searchByTitle(res http.ResponseWriter, req *http.Request) {
	title := req.PathValue("title")
	searchResult, err := s.SearchQuery(title)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	utils.JsonEncoder(res, searchResult)
}

func (s *Server) StartServer() {
	http.HandleFunc("/search/{title}", s.searchByTitle)
	http.ListenAndServe(":8080", nil)
}