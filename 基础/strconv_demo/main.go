package main

import (
	"fmt"
	"strconv"
)

func main() {
	//将字符串转换成数字
	str := "10000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return
	}
	fmt.Println(ret1)
	retint, _ := strconv.Atoi(str)
	fmt.Println(retint)

	//将数字转换成字符串
	s := int32(97)
	ret3 := fmt.Sprintf("%d", s)
	fmt.Printf("%#v", ret3)
}
