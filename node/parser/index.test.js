const parser = require("./index")

test('expect no hrefs from empty html', () => {
    expect(parser.ExtractLinksFromHTML('')).toStrictEqual([])
})

test('expect no links from link with no href', () => {
    expect(parser.ExtractLinksFromHTML('<html><a/></html>')).toStrictEqual([])
})

test('expect no hrefs from invalid html', () => {
    expect(parser.ExtractLinksFromHTML('<html><a href="google.com"')).toStrictEqual([])
})

test('expect href from valid self-closing html', () => {
    expect(parser.ExtractLinksFromHTML('<a href="google.com" />')).toStrictEqual(["google.com"])
})

test('expect href from valid html', () => {
    expect(parser.ExtractLinksFromHTML('<html><a href="google.com"><a/></html>')).toStrictEqual(["google.com"])
})

test('expect href from nested html', () => {
    expect(parser.ExtractLinksFromHTML('<html><body><h1><a href="google.com"></a></h1></body></html>')).toStrictEqual(["google.com"])
})