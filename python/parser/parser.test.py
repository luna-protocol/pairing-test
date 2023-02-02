import unittest
import parser

class Testing(unittest.TestCase):

    def test_extract_links_from_HTML_empty_reader(self):
        links = parser.ExtractLinksFromHTML("")
        self.assertEqual(links, [])

    def test_extract_links_from_invalid_HTML_reader(self):
        links = parser.ExtractLinksFromHTML("<html><body")
        self.assertEqual(links, [])
        
    def test_extract_links_from_link_with_no_href_reader(self):
        links = parser.ExtractLinksFromHTML("<html><a/></html>")
        self.assertEqual(links, [])

    def test_extract_links_from_valid_self_closing_link_reader(self):
        links = parser.ExtractLinksFromHTML('<a href="google.com" />')
        self.assertEqual(links, ["google.com"])
        
    def test_extract_links_from_valid_link_reader(self):
        links = parser.ExtractLinksFromHTML('<html><a href="google.com"><a/></html>')
        self.assertEqual(links, ["google.com"])

    def test_extract_links_from_valid_nested_link_reader(self):
        links = parser.ExtractLinksFromHTML('<html><body><h1><a href="google.com"></a></h1></body></html>')
        self.assertEqual(links, ["google.com"])


if __name__ == '__main__':
    unittest.main()