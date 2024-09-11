package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func init() {

	// Set log level based on environment variables
	switch logLevel := strings.ToLower(os.Getenv("LOG_LEVEL")); logLevel {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.ErrorLevel)
	}

	flag.Parse()

}

func main() {

	if *projectId == "" {
		log.Println("The 'project-id' flag is required")
		flag.Usage()
		os.Exit(1)
	}

	if *collectionName == "" {
		log.Println("The 'collection-name' flag is required")
		flag.Usage()
		os.Exit(1)
	}

	if *conversationName == "" {
		log.Println("The 'conversation-name' flag is required")
		flag.Usage()
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Missing input argument")
	}
	input := flag.Args()[0]

	apiKey := fetchAPIKey(projectId)
	payload := createRequestPayload(&input)
	response := sendRequest(payload, apiKey)

	// fmt.Println("Response: ")
	// fmt.Println(response)

	var chatCompletion ChatCompletion
	err = json.Unmarshal([]byte(response), &chatCompletion)
	if err != nil {
		log.Fatalf("Error decoding the JSON response: %v", err)
	}

	// Print the JSON response
	fmt.Println(*chatCompletion.Choices[0].Message.Content)
}
