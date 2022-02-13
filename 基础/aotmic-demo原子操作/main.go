package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var x int64

func add() {
	atomic.AddInt64(&x, 3) //x需传指针地址去修改
	defer wg.Done()
}

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
