package handlers

import (
	"art-prompt-api/db"
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
	prompt := "Generate a " + strconv.Itoa(number) + "words art prompt to draw. " + "Related to " + nouns[picked_noun] + ". " + "Return only one prompt and do not include the word draw in the prompt. Do not use the word generate or create in the prompt."

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

func formatGeneratedPrompt(prompt string) string {
	var formatted_prompt string
	removed_semicolon := strings.ReplaceAll(prompt, "\"", "")

	if strings.Contains(removed_semicolon, "\n\n") {
		parts := strings.Split(removed_semicolon, "\n\n")
		formatted_prompt = parts[1]
	} else {
		formatted_prompt = removed_semicolon
	}

	return formatted_prompt
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

	var prompt_message string
	//todo get userID from request
	user_key := "123" + "_prompt"
	does_user_prompt_exists, err := db.DoesKeyExists(user_key)
	if err != nil {
		fmt.Printf("Failed to check if user prompt exists: %v\n", err)
	}

	if does_user_prompt_exists {
		user_prompt, err := db.GetValue(user_key)
		if err != nil {
			fmt.Printf("Failed to get user cached prompt: %v\n", err)
		}

		request.Messages[0].Content = user_prompt
		prompt_message = user_prompt
	} else {
		ollama_response, err := sendOllamaRequest(ollama_POST_url, request)
		if err != nil {
			fmt.Printf("Failed to get art prompt: %v\n", err)
			http.Error(w, "Failed to get art prompt", http.StatusInternalServerError)
			return
		}

		model_message := formatGeneratedPrompt(ollama_response.Message.Content)
		prompt_message = model_message

		//todo set cahche to last until the next
		err = db.SetValue(user_key, model_message)
		if err != nil {
			fmt.Printf("Failed to cache user prompt: %v\n", err)
		}
	}

	response := map[string]string{"art_prompt": prompt_message}
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
