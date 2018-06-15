package main

import (
	"fmt"
	"sync"
)

type FetchedURL struct {
	urls map[string]bool
	mux  sync.Mutex
}

func (f *FetchedURL) Add(url string) {
	f.mux.Lock()
	f.urls[url] = true
	f.mux.Unlock()
}

func (f *FetchedURL) Exists(url string) bool {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.urls[url]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, fetched *FetchedURL) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	if fetched.Exists(url) {
		return
	}

	if depth <= 0 {
		return
	}

	ch := make(chan []string)

	go Fetch(url, ch)
	fetched.Add(url)

	urls := <-ch

	for _, u := range urls {
		Crawl(u, depth-1, fetcher, fetched)
	}
	return
}

func Fetch(url string, ch chan []string) {
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		close(ch)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	ch <- urls

}

func main() {
	fetched := &FetchedURL{urls: make(map[string]bool)}

	Crawl("https://golang.org/", 4, fetcher, fetched)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	// mapにkeyがあるか確認
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
