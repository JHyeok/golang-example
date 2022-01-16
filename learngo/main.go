package main

import (
	"fmt"

	"github.com/JHyeok/golang-example/learngo/banking"
)

func main() {
	// account := banking.Account{Owner: "jaehyeok", Balance: 1000}
	account := banking.Account{Owner: "JaeHyeok", Balance: 1000}
	fmt.Println(account)
}
