package concurrency

// WebsiteChecker 負責確認 url 的狀況
// 利用DI方便mock
type WebsiteChecker func(string) bool

type result struct {
	url   string
	alive bool
}


// CheckWebsites 確認 url 的狀態
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultCh := make(chan result)
	for _, url := range urls {
		go func()  {
			resultCh <- result{
				url: url,
				alive: wc(url),
			}
		}()
		
	}

	for range len(urls) {
		r := <- resultCh
		results[r.url] = r.alive
	}
	return results
}
