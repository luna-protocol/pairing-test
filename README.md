# Pairing Test

## The Problem

We'll be working together to write a simple web crawler in Golang. Don't worry if you've never written any Go before, the purpose of this exercise is to see how we work together and communicate! As such we care less about getting the code working and more about communication and building a shared understanding, so the more you explain what you're thinking the better. Just treat it like we're working together as colleagues.

The crawler should:

- Start with a given base domain
- Make a HTTP Get request to that domain to retrieve the HTML
- Parse the HTML to find links within the page (we've written this code for you).
- Recursively visit all of the found links, **but only if they are within the same domain**
- Print a list of all found links to the output

## Running Things

- `go run ./main.go` to run the program
- `go test ./...` to run any tests

## Tips

- Use `https://crawler-test.com` as the base domain
- Feel free to rearrange as you like! The only requirement is the CLI interface.
- Whether we use concurrency is up to you! But explain your reasoning.
