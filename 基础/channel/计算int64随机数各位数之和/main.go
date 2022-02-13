/*
1、开启一个goroutine循环生成int64随机数，发送到jobchan
2、开启24个goroutine从jobchan中取出随机数计算各位数的和，将结果发送给resultchan
3、主goroutine从resultchan取出结果并打印到终端输出
*/
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var jobchan = make(chan *job, 100)
var resultchan = make(chan *result, 100)
var wg sync.WaitGroup

func getint64(num chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newjob := &job{
			value: x,
		}
		num <- newjob
		time.Sleep(time.Millisecond * 500)
	}
}

func retint64(num <-chan *job, resultchan chan *result) {
	defer wg.Done()
	for {
		job := <-num
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newresult := &result{
			job: job,
			sum: sum,
		}
		resultchan <- newresult
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(1)
	go getint64(jobchan)
	wg.Add(24)
	for i := 0; i < 4; i++ {
		go retint64(jobchan, resultchan)
	}
	for result := range resultchan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()

}
