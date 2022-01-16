package main

import (
	"fmt"
	"net/http"
	"sync"
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

	// WaitGroup 생성
	wg := sync.WaitGroup{}

	for _, u := range urls {
		// WaitGroup에 대기 중인 고루틴 개수 추가
		wg.Add(1)
		go func(url string) {
			// 고루틴 종료 시 Done() 처리 - 대기 중인 고루틴의 수행이 종료되었다는 것을 알림
			defer wg.Done()
			checkUrl(url)
		}(u)
	}

	// 모든 고루틴이 종료될 때까지 대기
	wg.Wait()
}
