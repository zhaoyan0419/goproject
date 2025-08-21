package main

import "fmt"

func main() {
	addResult := func(num1 int, num2 int) int {
		return num2 + num1
	}(100, 200)
	fmt.Println("sunResult", addResult)

	sub := func(num1 int, num2 int) int {
		return num1 - num2
	}
	subResult := sub(200, 100)

	fmt.Println("sunResult", subResult)

}
