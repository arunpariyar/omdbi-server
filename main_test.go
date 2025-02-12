package main

import (
	"strconv"
	"testing"

	"github.com/arunpariyar/omdbi-server/server"
	"github.com/arunpariyar/omdbi-server/utils"
)

type tableTest struct {
	data   string
	result int
}

func setupTestServer() *server.Server {
	config := utils.GetEnv()
	testServer := server.NewServer(config["apiKey"]); 
	return testServer 
}

func TestSearchResults(t *testing.T) {
	
	testServer := setupTestServer()
	// add tableTest
	tests := []tableTest{
		{data: "honey", result: 638},
		{data: "juno", result: 52},
		{data: "moon", result: 2508},
	}

	for _, test := range tests {
		searchResult, err := testServer.SearchQuery(test.data)

		if err != nil {
			t.Errorf("Error converting string to int: %s", err)
		}

		amt, _ := strconv.Atoi(searchResult.TotalResults)

		if amt != test.result {
			t.Errorf("Expected 68 but got %s", searchResult.TotalResults)
		}
	}

}

