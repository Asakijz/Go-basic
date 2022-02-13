package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime caller() failed")
		return

	}
	funcname := runtime.FuncForPC(pc).Name()
	fmt.Println(funcname)
	// fmt.Println(file)
	fmt.Println(path.Base(file)) //获取最终路径
	fmt.Println(line)
}
