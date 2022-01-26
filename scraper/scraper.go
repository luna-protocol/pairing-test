package scraper

import "io"

type Scraper struct {
	BaseURL string
	Output  io.Writer
}

// Run should run the web scraper for the given domain
func (s *Scraper) Run() error {
	return nil
}
