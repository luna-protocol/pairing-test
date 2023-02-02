package main

import (
	"flag"
	"log"
	"os"

	"github.com/luna-protocol/pairing-test/scraper"
)

func main() {
	baseURL := flag.String("target", "https://crawler-test.com", "the initial url to scrape")
	flag.Parse()

	output := os.Stdout

	if err := scraper.Run(*baseURL, output); err != nil {
		log.Fatalf("Failed to run scraper: %v", err)
	}
}
