package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "https://api.debank.com/quest/list?limit=50&status=hot"
	client := &http.Client{}

	for {
		fmt.Println("start...")
		// Get the current timestamp
		// currentTimestamp := strconv.FormatInt(time.Now().Unix(), 10)

		req, err := http.NewRequest("GET", url, nil)
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
		req.Header.Set("x-api-nonce", "n_4FPV5vwEkS33pgKSQEmT0f3xcA9AGgC298C9cb64")
		req.Header.Set("x-api-sign", "d796dad8beba6805f9f4b6d7d1f62a2b11feb30173becabe4cac20247432f0e5")
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
