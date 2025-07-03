package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	MODEL    = "llama3-70b-8192"
	GROQ_URL = "https://api.groq.com/openai/v1/chat/completions"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model      string        `json:"model"`
	Messages   []Message     `json:"messages"`
	ToolChoice string        `json:"tool_choice,omitempty"`
	Tools      []interface{} `json:"tools,omitempty"`
}

type FunctionCall struct {
	Name string `json:"name"`
	Args string `json:"args,omitempty"`
}

type ToolCall struct {
	ID           string        `json:"id"`
	Type         string        `json:"type"`
	FunctionCall *FunctionCall `json:"function_call,omitempty"`
}

type Response struct {
	Reason    string     `json:"reason"`
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
	Message   Message    `json:"message,omitempty"`
}

var fun = []interface{}{
	map[string]interface{}{
		"type": "function",
		"function": map[string]interface{}{
			"name":        "get_files_info",
			"description": "List files / dirs in a directory",
			"parameters": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"path": map[string]interface{}{
						"type":        "string",
						"description": "the path to the directory",
					},
				},
				"required": []string{"path"},
			},
		},
	},
	map[string]interface{}{
		// TODO: Finish the function schemas for the AI
	}
}

func GenerateResponse(prompt string, verbose bool) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("error loading .env file: %v\n", err)
	}
	apiKey := os.Getenv("API_KEY")

	messages := []Message{
		{Role: "system", Content: "You are a helpful assistant."},
		{Role: "user", Content: prompt},
	}

	return resp, nil
}
