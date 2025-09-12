package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sigs.k8s.io/yaml"
)

func main() {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) != 4 {
		log.Println("参数输入有误，请输入3个参数。分别是yaml文件名和要查询的是volumes还是containers，具体字段名。 例如：")
		log.Printf("%s arg1 arg2 arg3", filepath.Base(execPath))
		return
	}
	filename := os.Args[1]
	arr := os.Args[2]
	fieldname := os.Args[3]
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var pod map[string]interface{}
	err = yaml.Unmarshal(data, &pod)
	if err != nil {
		log.Fatal(err)
	}
	spec := pod["spec"].(map[string]interface{})
	containers, ok := spec[arr].([]interface{})
	if !ok {
		fmt.Printf("arg2:%s Not Fount", arr)
		return
	}

	for k, v := range containers {
		c := v.(map[string]interface{})
		if c["name"] == fieldname {
			fmt.Println(k)
			return
		}
	}
	fmt.Printf("您输入的字段名：%s在本yaml中的%s中没有\n", fieldname, arr)
}
