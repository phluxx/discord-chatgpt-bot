package chatgpt

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "discord-chatgpt-bot/config"
)

const chatgptApiURL = "https://api.openai.com/v1/completions"

type ChatGPTRequest struct {
    Model     string `json:"model"`
    Prompt    string `json:"prompt"`
    MaxTokens int    `json:"max_tokens"`
}

type ChatGPTResponse struct {
    Choices []struct {
        Text string `json:"text"`
    } `json:"choices"`
}

func GetResponse(prompt string) (string, error) {
    reqBody, _ := json.Marshal(ChatGPTRequest{
        Model:     "gpt-4o-mini", 
        Prompt:    prompt,
        MaxTokens: 150,
    })

    req, err := http.NewRequest("POST", chatgptApiURL, bytes.NewBuffer(reqBody))
    if err != nil {
        return "", fmt.Errorf("failed to create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer " + config.ChatGPTAPIKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("request failed: %v", err)
    }
    defer resp.Body.Close()

    log.Printf("ChatGPT API response status: %s", resp.Status)
    respBody, _ := ioutil.ReadAll(resp.Body)
    log.Printf("ChatGPT API raw response body: %s", string(respBody))

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("non-200 status code from ChatGPT API: %s", resp.Status)
    }

    var chatResp ChatGPTResponse
    if err := json.Unmarshal(respBody, &chatResp); err != nil {
        return "", fmt.Errorf("failed to decode response: %v", err)
    }

    if len(chatResp.Choices) == 0 {
        return "", fmt.Errorf("no response from ChatGPT")
    }

    return chatResp.Choices[0].Text, nil
}

