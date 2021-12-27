package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkUrl(url string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "is down!")
		return
	}

	fmt.Println(url, "is up!")
}

func main() {
	urls := []string{
		"https://github.com/jhyeok",
		"https://jhyeok.com",
		"https://djqtsmstkdlxm.co",
	}

	for _, url := range urls {
		go checkUrl(url)
	}

	time.Sleep(5 * time.Second)
}
