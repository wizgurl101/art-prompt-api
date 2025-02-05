package handlers

import (
	models "art-prompt-api/models"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func generatePrompt() string {
	var nouns [8]string = [8]string{"video games", "sci fi", "animals", "nature", "anatomy", "plants", "space", "fantasy"}
	nouns_min := 0
	nouns_max := len(nouns) - 1
	picked_noun := rand.Intn(nouns_max-nouns_min+1) + nouns_min

	min := 2
	max := 4
	number := rand.Intn(max-min+1) + min
	prompt := "Generate a " + strconv.Itoa(number) + "words art prompt to draw. " + "Related to " + nouns[picked_noun] + ". " + "Return only one prompt and do not include the word draw in the prompt."

	return prompt
}

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
	prompt := generatePrompt()
	ollama_url := os.Getenv("OLLAMA_URL")
	ollama_POST_url := ollama_url + "/api/chat"

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
		fmt.Printf("Failed to get art prompt: %v\n", err)
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
