const cheerio = require('cheerio');

// ExtractLinksFromHTML expects the body to be a HTML string
function ExtractLinksFromHTML(body) {
    const $ = cheerio.load(body)
    const linkObjects = $('a'); // get all hyperlinks

    const links = [];

    linkObjects.each((_, element) => {
        const ref = $(element).attr('href')
        if (!!ref) {
            links.push(ref)
        }
    })

    console.log(links)
    return links
}

module.exports = {
    ExtractLinksFromHTML
}