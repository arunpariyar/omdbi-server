package main

import (
	"log"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

type tableTest struct {
	data   string
	result int
}
func TestSearchResults(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// add tableTest
	tests := []tableTest{
		{data: "honey", result: 638},
		{data: "juno", result: 52},
		{data: "moon", result: 2508},
	}

	for _, test := range tests {
		searchResult, err := SearchQuery(test.data)

		if err != nil {
			t.Errorf("Error converting string to int: %s", err)
		}

		amt, _ := strconv.Atoi(searchResult.TotalResults)

		if amt != test.result {
			t.Errorf("Expected 68 but got %s", searchResult.TotalResults)
		}
	}

}

