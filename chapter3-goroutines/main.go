package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [5]string{"jaehyeok", "kim", "test", "korea", "seoul"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
	// 만약 main 함수가 끝나면 goroutines은 무의미해진다.
	// result := <-c
	// fmt.Println("Waiting for messages")
	// 채널로부터 메시지를 가져오는 것
	// fmt.Println("Received this message: ", <-c)
	// fmt.Println("Received this message: ", <-c)
	// Goroutines은 메인 함수가 실행되어 있는 동안 이다. 이 5초가 끝나면 종료
	// time.Sleep(time.Second * 5)
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
