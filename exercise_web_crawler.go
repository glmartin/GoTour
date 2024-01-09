package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// SafeUrlMap is safe to use concurrently.
type SafeUrlMap struct {
	m sync.Mutex
	v map[string]string
}

func (u *SafeUrlMap) add(url, body string) {
	u.m.Lock()

	defer u.m.Unlock()
	if _, ok := u.v[url]; !ok {
		u.v[url] = body
	}
}

func (u *SafeUrlMap) getValue(url string) (string, bool) {
	u.m.Lock()
	defer u.m.Unlock()
	if body, ok := u.v[url]; ok {
		return body, true
	}
	return "", false
}

// Cache for the URLs
var urlMap SafeUrlMap = SafeUrlMap{}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan string) {
	if depth <= 0 {
		ch <- url
		return
	}

	if _, ok := urlMap.getValue(url); ok {
		// url already exists
		ch <- url
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)

		// save as empty if the URL is not found
		urlMap.add(url, "")
		ch <- url
		return
	}
	urlMap.add(url, body)
	fmt.Printf("found: %s %q\n", url, body)

	subCh := make(chan string)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, subCh)
	}

	for i := 0; i < len(urls); i++ {
		<-subCh
	}
	ch <- url

	return
}

func main() {
	urlMap.v = make(map[string]string)

	ch := make(chan string)
	go Crawl("https://golang.org/", 4, fetcher, ch)

	// wait to receive from the channel
	<-ch

	// Final solution should be :
	// https://golang.org/ "The Go Programming Language"
	// https://golang.org/pkg/ "Packages"
	// https://golang.org/cmd/
	// https://golang.org/pkg/fmt/ "Package fmt"
	// https://golang.org/pkg/os/ "Package os"
	for url, body := range urlMap.v {
		fmt.Printf("%s : %s\n", url, body)
	}
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
