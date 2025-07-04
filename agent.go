package gobootai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	FunctionCall FunctionCall `json:"function_call,omitempty"`
}

type ResponseChoice struct {
	FinishReason string     `json:"finish_reason"`
	ToolCalls    []ToolCall `json:"tool_calls,omitempty"`
	Message      Message    `json:"message,omitempty"`
}

type Response struct {
	Usage   map[string]int   `json:"usage"`
	Choices []ResponseChoice `json:"choices"`
}

var funs = []interface{}{
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
	},
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

	reqBody := Request{
		Model:      MODEL,
		Messages:   messages,
		Tools:      funs,
		ToolChoice: "auto",
	}

	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v\n", err)
	}

	// now to create a post request to the api
	resp, err := ApiRequest("POST", reqJSON, apiKey)
	if err != nil {
		return "", fmt.Errorf("error making the API request: %v\n", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var apiResp Response
	if err := json.Unmarshal(body, &apiResp); err != nil {
		fmt.Printf("Raw response: %s\n", string(body))
		return "", fmt.Errorf("error unmarshaling response: %v\n", err)
	}

	choice := apiResp.Choices[0]
	if choice.FinishReason == "tool_calls" && len(choice.ToolCalls) > 0 {
		tool := choice.ToolCalls[0]
		return fmt.Sprintf("Ai agent called func: %s\nWith args: %s\n", tool.FunctionCall.Name, tool.FunctionCall.Args), nil
	}

	return choice.Message.Content, nil
}

func ApiRequest(method string, body []byte, apiKey string) (*http.Response, error) {
	req, err := http.NewRequest(method, GROQ_URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s\n", resp.Status)
	}

	return resp, nil
}
