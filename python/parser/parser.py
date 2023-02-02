from bs4 import BeautifulSoup

def ExtractLinksFromHTML(body):
    soup = BeautifulSoup(body, 'html.parser')
    
    links = []
    
    for link in soup.find_all('a'):
        path = link.get('href')
        if path:
            links.append(path)
        
    return links