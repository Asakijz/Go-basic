package main

import (
	"fmt"
	"sync"
)

var a chan int
var wg sync.WaitGroup

func main() {
	fmt.Println(a) //nil
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-a
		fmt.Println("后台goroutine从通道中读取到了", x)
	}()
	a = make(chan int) //不带缓冲区通道的初始化
	a <- 10            //不开辟空间-->蚌埠住了
	fmt.Println("10发送到了通道中")
	a = make(chan int, 16) //带缓冲区通道的初始化--->如果后面数值过大建议将原类型改为指针类型
	fmt.Println(a)
	wg.Wait()
}
