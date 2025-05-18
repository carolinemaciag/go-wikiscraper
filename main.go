package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	start := time.Now()

	// List of Wikipedia URLs to scrape
	urls := []string{
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Android_(robot)",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
	}

	var wg sync.WaitGroup
	mu := &sync.Mutex{}
	scrapedData := make(map[string]string)

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	// Extract text content from paragraphs
	c.OnHTML("p", func(e *colly.HTMLElement) {
		url := e.Request.URL.String()
		mu.Lock()
		scrapedData[url] += e.Text + "\n"
		mu.Unlock()
	})

	// Log when scraping each page finishes
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scraped:", r.Request.URL)
	})

	// Log errors
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request to %s failed with response %v: %v", r.Request.URL, r, err)
	})

	// Start scraping in goroutines for each URL
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			err := c.Visit(u)
			if err != nil {
				log.Printf("Visit error for %s: %v", u, err)
			}
		}(url)
	}

	wg.Wait() // Wait for all scraping goroutines to finish

	// Define struct for JSON output
	type PageData struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	}

	// Create the JSON lines output file
	file, err := os.Create("scraped_output.jsonl")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	// Write each scraped page's data as a JSON line
	for url, content := range scrapedData {
		entry := PageData{
			URL:  url,
			Text: strings.TrimSpace(content),
		}
		jsonBytes, err := json.Marshal(entry)
		if err != nil {
			log.Printf("Error marshaling JSON for URL %s: %v", url, err)
			continue
		}
		file.Write(jsonBytes)
		file.Write([]byte("\n"))
	}

	elapsed := time.Since(start)
	fmt.Printf("\nScraping completed in %s\n", elapsed)
	fmt.Println("Scraped data saved to scraped_output.jsonl")
}
