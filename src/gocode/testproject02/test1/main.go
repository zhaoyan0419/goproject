package main

func main() {
	c1 := make(chan int, 5)
	c1 <- 10
	c1 <- 20
	c1 <- 30
	c1 <- 40
	c1 <- 50
	c1 <- 60
}
