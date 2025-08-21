package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "192.168.1.100 - - [10/Oct/2023:14:32:02 +0800] \"POST /api/login HTTP/1.1\" 200 567"
	s1 := strings.Fields(str1)
	for _, v := range s1 {
		fmt.Printf("%v\n", v)
	}
}
