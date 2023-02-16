# Pairing Test

## The Problem

Our largest competitor just launched a new website. We need to support our business analysts to understand the structure of the website, so they've asked for a list of all the pages on the website.

## The Ask

We'll be working together to write a simple web crawler in a language of your choice. The purpose of this exercise is to see how we work together and how you approach a problem! As such we care less about getting the code working and more about communication and building a shared understanding, so the more you explain what you're thinking the better. Just treat it like we're working together as colleagues.

The crawler should:

1. Start with from a single URL
2. Make a HTTP Get request to that URL to retrieve the HTML
3. Parse the HTML to find links within the page (we've written this code for you).
4. Keep repeating steps 2. and 3. for all of the found links **but only if they are within the same domain** until you've visited every page on the website. 
    - For example: ['/home', 'https://crawler-test.com/v1/done', 'https://google.com'] you'd visit the first two links. 
5. Print a list of all the links within the site to the terminal.

## Tips

- Use `https://crawler-test.com` as the base domain
- Feel free to rearrange as you like! The only requirement is the CLI interface.
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
