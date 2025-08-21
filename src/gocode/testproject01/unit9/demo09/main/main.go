package main

import (
	"fmt"
	"github.com/src/gocode/testproject01/unit9/demo09/model"
)

func main() {
	c1 := &model.Cat{}
	c1.Animal.Age = 10
	c1.Animal.Weight = 15
	c1.Animal.Shout()
	c1.Animal.ShowInfo()
	c1.Scrath()
	fmt.Println("------------------------------------")
	// 引用字段时候可以缩写
	c1.Age = 20
	c1.Weight = 25
	c1.Shout()
	c1.ShowInfo()
	c1.Scrath()
}
