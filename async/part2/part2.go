package main

import (
	"fmt"
	"net/http"
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
}
