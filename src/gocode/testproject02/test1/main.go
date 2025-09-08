package main

import "log"

func main() {
	var c1 chan int

	c1 = make(chan int, 100)
	close(c1)
	if c1 == nil {
		log.Println("c1 为nil")
	}
	log.Println(c1)

}
