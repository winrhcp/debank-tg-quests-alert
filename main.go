package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return
	}

	apiURL := os.Getenv("API_URL")
	apiNonce := os.Getenv("API_NONCE")
	apiSign := os.Getenv("API_SIGN")
	retryInterval := 1 * time.Minute
	seenQuestIDs, err := initSeenQuest(apiURL, apiNonce, apiSign)
	if err != nil {
		log.Fatal(err)
		return
	}

	round := 0
	for {
		fmt.Printf("Running round %d...\n", round)
		// Get the current timestamp
		// currentTimestamp := strconv.FormatInt(time.Now().Unix(), 10)

		body, err := fetchQuestData(apiURL, apiNonce, apiSign)
		if err != nil {
			log.Printf("Request failed: %v", err)
			time.Sleep(retryInterval)
			continue
		}

		var questResponse QuestResponse
		err = json.Unmarshal(body, &questResponse)
		if err != nil {
			log.Printf("Unmarshal failed: %v", err)
			time.Sleep(retryInterval)
			continue
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

		time.Sleep(1 * time.Minute)
		round++
	}
}
