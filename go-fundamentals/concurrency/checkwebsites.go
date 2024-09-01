package concurrency

type WebsiteChecker func(string) bool

type result struct {
	url     string
	success bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan result)

	for _, url := range urls {
		url := url
		go func() {
			resultChan <- result{
				success: wc(url),
				url:     url,
			}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChan
		results[r.url] = r.success
	}

	return results
}
