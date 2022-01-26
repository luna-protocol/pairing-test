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

	scraper := scraper.Scraper{
		BaseURL: *baseURL,
		Output:  output,
	}

	if err := scraper.Run(); err != nil {
		log.Fatal(err)
	}
}
