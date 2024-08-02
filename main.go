package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	client := &http.Client{}

	for {
		fmt.Println("start...")
		// Get the current timestamp
		// currentTimestamp := strconv.FormatInt(time.Now().Unix(), 10)

		// Load environment variables from .env file
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		apiURL := os.Getenv("API_URL")
		apiNonce := os.Getenv("API_NONCE")
		apiSign := os.Getenv("API_SIGN")

		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			log.Fatal(err)
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
		req.Header.Set("x-api-ts", "1722583294")
		req.Header.Set("x-api-ver", "v2")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close() // Ensure the body is closed to prevent resource leaks
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var questResponse QuestResponse
		err = json.Unmarshal(body, &questResponse)
		if err != nil {
			log.Fatal(err)
		}

		// Print the quests
		// fmt.Println(questResponse)
		for _, quest := range questResponse.Data.Quests {
			fmt.Println(quest.Article.ID)
		}

		time.Sleep(1 * time.Minute)
	}
}
