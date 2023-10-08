package main

// OpenAI ChatCompletion structs
type ChatCompletion struct {
	Id      string   `json:"id,omitempty" firestore:"id,omitempty"`
	Object  *string  `json:"object,omitempty" firestore:"object,omitempty"`
	Created *int64   `json:"created,omitempty" firestore:"created,omitempty"`
	Model   *string  `json:"model,omitempty" firestore:"model,omitempty"`
	Choices []Choice `json:"choices,omitempty" firestore:"choices,omitempty"`
	Usage   *Usage   `json:"usage,omitempty" firestore:"usage,omitempty"`
}

type Choice struct {
	Index        *int           `json:"index,omitempty" firestore:"index,omitempty"`
	Message      *OpenAiMessage `json:"message,omitempty" firestore:"message,omitempty"`
	FinishReason *string        `json:"finish_reason,omitempty" firestore:"finish_reason,omitempty"`
}

type OpenAiMessage struct {
	Role    *string `json:"role,omitempty" firestore:"role,omitempty"`
	Content *string `json:"content,omitempty" firestore:"content,omitempty"`
}

type Usage struct {
	PromptTokens     *int `json:"prompt_tokens,omitempty" firestore:"prompt_tokens,omitempty"`
	CompletionTokens *int `json:"completion_tokens,omitempty" firestore:"completion_tokens,omitempty"`
	TotalTokens      *int `json:"total_tokens,omitempty" firestore:"total_tokens,omitempty"`
}
