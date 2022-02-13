package main

import "fmt"

func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println(t)
	case bool:
		fmt.Println(t)
	case int:
		fmt.Println(t)
	}

}

func main() {
	assign2(100)
	assign2(true)
	assign2("张三")

}
