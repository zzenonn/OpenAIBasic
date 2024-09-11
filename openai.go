package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

const (
	OPENAI_URL  = "https://api.openai.com/v1/chat/completions"
	TEMPERATURE = 0.3
)

func getSecret(secretName string) (string, error) {
	ctx := context.Background()

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	secretRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretName,
	}

	result, err := client.AccessSecretVersion(ctx, secretRequest)
	if err != nil {
		return "", err
	}

	return string(result.Payload.Data), nil
}

func fetchAPIKey(projectName *string) string {
	openAIKeySecretName := fmt.Sprintf("projects/%s/secrets/OpenAIAPIKey/versions/latest", *projectName)
	apiKey, err := getSecret(openAIKeySecretName)
	if err != nil {
		log.Fatalf("Failed to fetch the secret using the name %s. Error: %s", openAIKeySecretName, err.Error())
	}
	return apiKey
}

func createRequestPayload(input *string) []byte {

	modelVal := "chatgpt-4o-latest"
	roleSystem := "system"
	systemContent := ``
	roleUser := "user"

	data := map[string]interface{}{
		"model":       &modelVal,
		"temperature": TEMPERATURE,
		"messages": []*OpenAiMessage{
			{
				Role:    &roleSystem,
				Content: &systemContent,
			},
			{
				Role:    &roleUser,
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	return string(body)
}
