package handlers

import (
	models "art-prompt-api/models"
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

func sendOllamaRequest(url string, request models.OllamaRequest) (models.OllamaResponse, error) {
	var response models.OllamaResponse

	jsonData, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return response, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetArtPrompt(w http.ResponseWriter, r *http.Request) {
	prompt := "Generate a 3 words art prompt to draw. Give only one prompt."
	ollama_POST_url := "http://localhost:11434/api/chat"

	request := models.OllamaRequest{
		Model: "llama3.2:1b",
		Messages: []models.Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Stream: false,
	}

	ollama_response, err := sendOllamaRequest(ollama_POST_url, request)
	if err != nil {
		http.Error(w, "Failed to get art prompt", http.StatusInternalServerError)
		return
	}

	model_message := strings.ReplaceAll(ollama_response.Message.Content, "\"", "")
	response := map[string]string{"art_prompt": model_message}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ArtPromptHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetArtPrompt(w, r)
	default:
		http.Error(w, "404 Not Found", http.StatusMethodNotAllowed)
	}
}
