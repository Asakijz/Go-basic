package main

import (
	"bufio"
	"fmt"
	"os"
)

func usebufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容：")
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n", s)
}

func main() {
	usebufio()
}
