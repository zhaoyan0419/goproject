package main

import (
	"fmt"
	"github.com/src/gocode/testproject01/unit9/demo07/model1"
)

func main() {
	var s1 model1.Student = model1.Student{Name: "yanyan", Age: 28}
	fmt.Println(s1)
	p1 := model1.Teacher{"zzss"}
	fmt.Println(p1)
	p2 := model1.NewPersion("zzyy", 28)
	fmt.Println(*p2)
	fmt.Printf("my name is %v, i am %d years old\n", p2.Name, p2.Age)
	f1 := model1.NewmyFirstStruct("dsdsdsdsds")
	fmt.Println(f1.GetName())
}
