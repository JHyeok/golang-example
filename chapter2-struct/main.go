package main

import (
	"fmt"

	"github.com/JHyeok/golang-example/chapter2-struct/accounts"
)

func main() {
	account := accounts.NewAccount("jaehyeok")
	account.Deposit(10)
	// Go에서 error를 handle하는 방법
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
