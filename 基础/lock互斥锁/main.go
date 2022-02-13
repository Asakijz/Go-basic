package main

import (
	"fmt"
	"sync"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

//lock锁存在时，多个goroutine只能等待上一个进程结束才能继续执行
func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	defer wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
