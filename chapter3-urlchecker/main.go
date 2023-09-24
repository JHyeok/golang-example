package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	// "golang.org/x/net/context"
)

/* Channel을 사용하는 방식
type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("request failed")

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://jhyeok.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
*/

func main() {
	var wg sync.WaitGroup
	results := make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://jhyeok.com/",
	}

	for _, url := range urls {
		wg.Add(1)
		go hitURL(url, &wg, results)
	}

	wg.Wait()

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, wg *sync.WaitGroup, results map[string]string) {
	defer wg.Done()

	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	results[url] = status
}

/*
func consuming(scheduler chan string) {
	select {
	case <-scheduler:
		fmt.Println("이름을 입력받았습니다.")
	case <-time.After(5 * time.Second):
		fmt.Println("시간이 지났습니다.")
	}
}

func producing(scheduler chan string) {
	var name string
	fmt.Print("이름:")
	fmt.Scanln(&name)
	scheduler <- name
}

func main() {
	scheduler := make(chan string)
	go consuming(scheduler)
	go producing(scheduler)

	time.Sleep(100 * time.Second)
}
*/

type Datastore interface {
	Fetch(ctx context.Context, key string) (interface{}, error)
	Update(ctx context.Context, key string, f func(current interface{}) (interface{}, error)) (interface{}, error)
}
