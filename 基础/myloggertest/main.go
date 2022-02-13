package main

import (
	"mylogger"
	"sync"
)

var Wg sync.WaitGroup

//测试日志库

func main() {
	// log := mylogger.Newlog("error")
	log := mylogger.Newfilelogger("info", "./", "xx.log", 10*1024*1024)
	for {
		// Wg.Add(1)
		log.Debug("这是一条debug日志")
		log.Trace("这是一条trace日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条warning日志")
		id := 10010
		name := "alex"
		//error级别往下输出.err日志文件
		log.Error("这是一条error日志,id:%d,name:%s", id, name)
		log.Fatal("这是一条fatal日志")
		// time.Sleep(time.Second * 3)
		// Wg.Wait()
	}
}
