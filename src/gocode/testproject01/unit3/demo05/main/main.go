package main

import "fmt"

func getSum() func(int) int {
	sum := 0
	return func(i int) int {
		return sum + i
	}
}

func main() {
	sumFunc := getSum()
	fmt.Println(sumFunc(100))
}
