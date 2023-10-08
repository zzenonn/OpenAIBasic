package main

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/firestore"
	log "github.com/sirupsen/logrus"
)

type FirestoreDb struct {
	Client *firestore.Client
}

func NewDatabase(projectId string) (*FirestoreDb, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &FirestoreDb{
		Client: client,
	}, nil
}

// Convert ChatCompletion into a map
func StructToMap(input interface{}) map[string]interface{} {
	data, err := json.Marshal(input)
	if err != nil {
		log.Fatalf("Error marshalling the struct: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalf("Error unmarshalling into a map: %v", err)
	}

	return result
}

// func (db *FirestoreDb) GetConversation(conversationName string) (*Conversation, error) {
// 	ctx := context.Background()
// 	doc, err := db.Client.Collection("conversations").Doc(conversationName).Get(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var conversation Conversation
// 	err = doc.DataTo(&conversation)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &conversation, nil
// }
