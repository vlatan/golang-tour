package concurrency

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

type SafeMap struct {
	mu sync.Mutex
	v  map[string]struct{} // Using struct{} as value to store zero bytes
}

// lock the map, check if key exists and add it if it doesn't
func (sm *SafeMap) Seen(key string) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if _, ok := sm.v[key]; ok {
		return true
	}
	sm.v[key] = struct{}{}
	return false
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {

	// safe map to store visited urls
	var sm = SafeMap{v: make(map[string]struct{})}
	// wait group to wait all recursive goroutines to finish
	var wg sync.WaitGroup

	var recurse func(string, int)
	recurse = func(url string, depth int) {
		// on exit of the function decrement wait group by one
		defer wg.Done()

		// if depth zero or url seen, exit
		if depth <= 0 || sm.Seen(url) {
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		for _, u := range urls {
			wg.Add(1) // add one to wait group countyer before every goroutine starts
			go recurse(u, depth-1)
		}
	}

	wg.Add(1) // add one to wait group counter before every goroutine starts
	go recurse(url, depth)
	// wait/block untill the weight group counter is zero
	wg.Wait()
}
