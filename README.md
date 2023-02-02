# Pairing Test

## The Problem

Our largest competitor just launched a new website. We need to support our business analysts to understand the structure of the website, so they've asked for a list of all the pages on the website.

## The Ask

We'll be working together to write a simple web crawler in a language of your choice. The purpose of this exercise is to see how we work together and how you approach a problem! As such we care less about getting the code working and more about communication and building a shared understanding, so the more you explain what you're thinking the better. Just treat it like we're working together as colleagues.

The crawler should:

- Start with from a single URL
- Make a HTTP Get request to that URL to retrieve the HTML
- Parse the HTML to find links within the page (we've written this code for you).
- Recursively visit all of the found links, **but only if they are within the same domain**
- Print a list of all found links to the output stream

## Tips

- Use `https://crawler-test.com` as the base domain
- Feel free to rearrange as you like! The only requirement is the CLI interface.
- Whether we use concurrency is up to you! But explain your reasoning.
- Try and avoid libraries. `npm install webcrawler` is not a valid solution!

##  Running the repos

### Golang

```bash
cd go
go run main.go

# tests
go test ./...
```

### Node

```bash
cd node
node index.js

# tests 
npm t
```

### Python

```bash
cd python 
pip install -r requirements.txt
python main.py

# tests 
python parser/parser.test.py
```
