package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// fetchQuestData makes an HTTP GET request to the specified URL with headers and returns the response body.
func fetchQuestData(url, apiNonce, apiSign string) ([]byte, error) {
	client := &http.Client{}
	// currentTimestamp := strconv.FormatInt(time.Now().Unix(), 10)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en,th-TH;q=0.9,th;q=0.8")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("origin", "https://debank.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://debank.com/")
	req.Header.Set("sec-ch-ua", `"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("source", "web")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")
	req.Header.Set("x-api-nonce", apiNonce)
	req.Header.Set("x-api-sign", apiSign)
	req.Header.Set("x-api-ts", "1722569054")
	req.Header.Set("x-api-ver", "v2")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}

func initSeenQuest(url, apiNonce, apiSign, botToken, channelID string) (map[int]struct{}, error) {
	fmt.Println("init.....")
	seenQuestIDs := make(map[int]struct{})
	body, err := fetchQuestData(url, apiNonce, apiSign)
	if err != nil {
		return nil, err
	}

	var questResponse QuestResponse
	err = json.Unmarshal(body, &questResponse)
	if err != nil {
		return nil, err
	}

	for _, quest := range questResponse.Data.Quests {
		questID := quest.Article.ID
		_, seen := seenQuestIDs[questID]
		if !seen {
			// New quest found, print it
			fmt.Printf("New Quest: %+v\n", quest.Article.Quest.Name)
			fmt.Printf("Link: %s\n", createQuestURL(questID))
			seenQuestIDs[questID] = struct{}{}
		}
	}
	return seenQuestIDs, nil

}

func createQuestURL(id int) string {
	r := "91299"
	return fmt.Sprintf("https://debank.com/stream/%d?r=%s", id, r)
}

// SendMessage sends a message to a specified Telegram chat ID
func SendMessage(botToken, text, textButton, url, chatID string) error {

	// Create the API URL for sending a message
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	button := InlineKeyboardButton{
		Text: textButton,
		URL:  url,
	}

	inlineKeyboard := InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{
			{button},
		},
	}

	payload := map[string]interface{}{
		"chat_id":      chatID,
		"text":         text,
		"reply_markup": inlineKeyboard,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Send the request
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	return nil
}
