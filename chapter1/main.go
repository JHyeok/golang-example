package main

import (
	"fmt"
	"strings"
)

func multiply(x, y int) int {
	return x * y
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpperNaked(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	// return length, uppercase
	return
}

func lenAndUpperNakedDefer(name string) (length int, uppercase string) {
	// function이 값을 리턴하면 실행 defer
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	// return length, uppercase
	return
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	// variable expression
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrinkSwitch(age int) bool {
	/*
		// variable expression
		switch koreanAge := age + 2; koreanAge {
		case 10:
			return false
		case 18:
			return true
		}
	*/
	/*
		switch age {
		case 10:
			return false
		case 18:
			return true
		}
	*/
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	hyeok := person{name: "JHyeok", age: 20, favFood: favFood}
	// hyeok := person{"JHyeok", 20, favFood}
	fmt.Println(hyeok.name)

	jaehyeok := map[string]string{"name": "jaehyeok", "age": "20"}
	for _, value := range jaehyeok {
		fmt.Println(value)
	}
	fmt.Println(jaehyeok)

	// names := [5]string{"kim", "jae", "hyeok"}
	names := []string{"kim", "jae", "hyeok"}
	names = append(names, "test")
	fmt.Println(names)

	a := 2
	b := &a
	*b = 20
	// 메모리 주소: &, 메모리 주소에 뭐가 있는지 보려면: *
	fmt.Println(a)

	fmt.Println(canIDrinkSwitch(18))

	fmt.Println(canIDrink(16))

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	repeatMe("jae", "kim", "hyeok", "test")
	fmt.Println(multiply(2, 2))

	// _, upperName := lenAndUpperNakedDefer("jaehyeok")
	nameLength, upperName := lenAndUpperNakedDefer("jaehyeok")
	fmt.Println(nameLength, upperName)

	// var name string = "jaehyeok"
	// 축약형은 오로지 func 안에서만 가능
	name := "jaehyeok"
	name = "kim"
	fmt.Println(name)
}
