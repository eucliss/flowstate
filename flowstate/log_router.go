package flowstate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type LogRouter struct {
	Source      string
	Destination string
}

func NewLogRouter() *LogRouter {
	return &LogRouter{}
}

func (lr *LogRouter) Router() {
	logs, err := gatherLogs(lr.Source)
	if err != nil {
		fmt.Printf("failed to get logs: %v", err)
	}
	fmt.Printf("Received logs: %s\n", logs)

	routeLogs(logs, lr.Destination)

}

func gatherLogs(source string) (interface{}, error) {

	// Make GET request\

	// Make GET request
	resp, err := http.Get(source)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	jsonValue := map[string]interface{}{}
	err = json.Unmarshal(body, &jsonValue)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	// TODO: Process the logs from body
	fmt.Printf("Received logs: %s\n", jsonValue["result"])

	return jsonValue["result"], nil
}

func routeLogs(logs interface{}, destination string) {
	// POST API call to destination
	jsonValue, err := json.Marshal(logs)
	if err != nil {
		fmt.Printf("failed to marshal logs: %v", err)
	}

	err = godotenv.Load()
	if err != nil {
		fmt.Printf("failed to load .env file: %v", err)
	}
	// POST API call to destination
	req, err := http.NewRequest("POST", destination, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("failed to create request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	fmt.Printf("Username: %s\n", os.Getenv("USERNAME"))
	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("failed to post logs: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v", err)
		return
	}
	fmt.Printf("Response body: %s\n", string(body))

}

// func gatherLogs(source string) ([]byte, error) {

// 	// Make GET request\

// 	// Make GET request
// 	resp, err := http.Get(source)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get logs: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read response: %v", err)
// 	}

// 	return body, nil
// }

// func routeLogs(logs []byte, destination string) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Printf("failed to load .env file: %v", err)
// 	}
// 	// POST API call to destination
// 	req, err := http.NewRequest("POST", destination, bytes.NewBuffer(logs))
// 	if err != nil {
// 		fmt.Printf("failed to create request: %v", err)
// 		return
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	fmt.Printf("Username: %s\n", os.Getenv("USERNAME"))
// 	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
// 	// Send the request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Printf("failed to post logs: %v", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Printf("failed to read response body: %v", err)
// 		return
// 	}
// 	fmt.Printf("Response body: %s\n", string(body))

// }
