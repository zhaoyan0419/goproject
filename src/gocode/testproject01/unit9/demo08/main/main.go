package main

import (
	"fmt"
	model "github.com/src/gocode/testproject01/unit9/demo08/model1"
)

func main() {
	p1 := model.NewPerson("赵岩")
	p1.Setage(149)
	println(p1.GetAge())
	fmt.Println(*p1)
}
