package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//按行读取文件
func readfile() {
	fileobj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
	}
	defer fileobj.Close()
	reader := bufio.NewReader(fileobj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed,err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

//直接读取

func justread() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	fmt.Print(string(ret))
}

func main() {
	// readfile()
	justread()
}
