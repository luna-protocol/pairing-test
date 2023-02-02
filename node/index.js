const scraper = require("./scraper")


function run() {
    const baseURL = "https://crawler-test.com"
    try {
        scraper.Run(baseURL)
    } catch (err) {
        console.error("Failed to run scraper: ", err)
    }
}


run()