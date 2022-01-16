package main

import (
	"fmt"
	"net/http"
)

type urlStatus struct {
	url    string
	status bool
}

func checkUrl(url string, c chan urlStatus) {
	_, err := http.Get(url)

	if err != nil {
		// c 채널에 메시지 전송
		c <- urlStatus{url, false}
	} else {
		// c 채널에 메시지 전송
		c <- urlStatus{url, true}
	}
}

func main() {
	urls := []string{
		"https://github.com/jhyeok",
		"https://jhyeok.com",
		"https://djqtsmstkdlxm.co",
	}

	// urlStatus 채널 생성
	c := make(chan urlStatus)
	for _, url := range urls {
		// checkUrl을 고루틴으로 실행하고, c 채널을 매개변수로 넘김
		go checkUrl(url, c)
	}

	result := make([]urlStatus, len(urls))
	for i, _ := range result {
		// c 채널로부터 메시지 수신
		result[i] = <-c
		if result[i].status {
			fmt.Println(result[i].url, "is up!")
		} else {
			fmt.Println(result[i].url, "is down!")
		}
	}
}
