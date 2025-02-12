package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arunpariyar/omdbi-server/models"

	"github.com/joho/godotenv"
)

func GetEnv() map[string]string {
	LoadEnv()
	
	config := make(map[string]string)
	config["apiKey"] = os.Getenv("OMDB_API_KEY")
	return config
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func JsonDecoder(res *http.Response) (models.SearchResult, error) {
	result := models.SearchResult{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&result); err != nil {
		return result, fmt.Errorf("failed to decode json: %w", err)
	}

	return result, nil
}

func JsonEncoder(res http.ResponseWriter, data models.SearchResult) {
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(data); err != nil {
		log.Println(err)
	}
}
