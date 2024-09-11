package main

import (
	"flag"
)

var (
	projectId        = flag.String("project-id", "", "The id of the project (required)")
	collectionName   = flag.String("collection-name", "", "The Firestore collection name (required)")
	conversationName = flag.String("conversation-name", "", "The conversation name (required)")
	outputRaw        = flag.Bool("output-raw", false, "outputs the raw response from OpenAI")
)

