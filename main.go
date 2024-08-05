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
	// Replace with your Telegram Bot Token
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		fmt.Errorf("Telegram bot token not found in environment variables")
		return
	}
	channelID := os.Getenv("CHANNEL_ID")
	if channelID == "" {
		log.Fatal("CHANNEL_ID not set in environment")
	}
	retryInterval := 1 * time.Minute

	seenQuestIDs, err := initSeenQuest(apiURL, botToken, channelID)
	if err != nil {
		log.Fatal(err)
		return
	}

	round := 0
	for {
		fmt.Printf("Running round %d...\n", round)

		body, err := fetchQuestData(apiURL)
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
				// New quest found, send message to channel
				xp := quest.Article.Quest.UnitXP
				textButton := fmt.Sprintf("View Quest %d XP", xp)
				message := fmt.Sprintf("New Quest: %s\n", quest.Article.Quest.Name)
				if err := SendMessage(botToken, message, textButton, createQuestURL(questID), channelID); err != nil {
					log.Printf("Failed to send message: %v", err)
				}
				seenQuestIDs[questID] = struct{}{}
			}
		}

		time.Sleep(1 * time.Minute)
		round++
	}
}
