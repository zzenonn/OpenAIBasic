package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	OPENAI_URL  = "https://api.openai.com/v1/chat/completions"
	TEMPERATURE = 0.1
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func fetchAPIKey() string {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is not set")
	}
	return apiKey
}

func createRequestPayload(input string) []byte {
	data := map[string]interface{}{
		"model":       "gpt-4",
		"temperature": TEMPERATURE,
		"messages": []Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant.",
			},
			{
				Role:    "user",
				Content: input,
			},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to encode data to JSON: %v", err)
	}

	return jsonData
}

func sendRequest(data []byte, apiKey string) string {
	req, err := http.NewRequest("POST", OPENAI_URL, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make the request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	return string(body)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Missing input argument")
	}
	input := flag.Args()[0]

	apiKey := fetchAPIKey()
	payload := createRequestPayload(input)
	response := sendRequest(payload, apiKey)

	// Print the JSON response
	fmt.Println(response)
}
