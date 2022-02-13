package main

import "fmt"

type person struct {
	name  string
	age   int
	hobby []string
}

//结构体 type **** struct  可以批量声明

func main() {
	var p person
	p.name = "zhangsan"
	p.age = 18
	p.hobby = []string{"nan,nv"}
	fmt.Println(p)
}
