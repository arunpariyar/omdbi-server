// this is the main file for the omdbi-server
package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/arunpariyar/omdbi-server/models"
	"github.com/arunpariyar/omdbi-server/utils"
)

const URL string = "http://www.omdbapi.com/?apikey="

//CREATING CACHE
var omdbiCache = map[string]models.SearchResult{}

// searchQuery function to search for a movie by title
func SearchQuery(query string) (models.SearchResult, error) {

	if value, exists := omdbiCache[query]; exists {
		return value, nil
	}

	apiKey := os.Getenv("OMDB_API_KEY")
	if apiKey == "" {
		return models.SearchResult{}, fmt.Errorf("API key not found")
	}

	queryURL := URL + apiKey + "&s=" + query

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

	omdbiCache[query] = searchResult
	return searchResult, nil
}

// searchByTitle function to search for a movie by title
func searchByTitle(res http.ResponseWriter, req *http.Request) {
	title := req.PathValue("title")
	searchResult, err := SearchQuery(title)

	if err != nil {
		log.Println(err)
	}

	utils.JsonEncoder(res, searchResult)
}

func main() {
	utils.LoadEnv()
	http.HandleFunc("/search/{title}", searchByTitle)
	http.ListenAndServe(":8080", nil)
}
