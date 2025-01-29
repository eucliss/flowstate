package monitor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type OpenObserve struct {
	URL string
}

func (o *OpenObserve) Query(query Query) []interface{} {
	endpoint := o.URL + "/api/" + query.SourceType + "/_search"

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("failed to load .env file: %v", err)
	}

	initialQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"sql":        query.Query,
			"from":       0,
			"size":       query.Limit,
			"start_time": query.Start,
			"end_time":   query.End,
			"sql_mode":   "full",
		},
		"search_type": "ui",
	}
	jsonValue, err := json.Marshal(initialQuery)
	if err != nil {
		fmt.Printf("failed to marshal query: %v", err)
	}
	fmt.Printf("jsonValue: %s\n", string(jsonValue))

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("failed to create request: %v", err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("failed to post logs: %v", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v", err)
		return nil
	}
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	// fmt.Printf("result: %v\n", result)
	// fmt.Println("--------------------------------")

	// fmt.Printf("result hits: %v\n", result["hits"])
	// fmt.Println("--------------------------------")
	return result["hits"].([]interface{})
}
