package scraper

import (
	"fmt"
	"io"
)

// Run should run the web scraper for the given domain
// output can be written using output.Write([]byte("Hello, World!"))
func Run(baseURL string, output io.Writer) error {
	output.Write([]byte(fmt.Sprintf("Fetching base url: %s\n", baseURL)))
	return nil
}
