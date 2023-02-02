import scraper

def run():
    baseURL = "https://crawler-test.com"
    try:
        scraper.Run(baseURL)
    except Exception as e:
        print(f'Failed to run scraper: {e}')


if __name__ == '__main__':
    run()