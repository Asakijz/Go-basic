package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	for {
		fileobj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("open file failed,err:%v", err)
			return
		}
		log.SetOutput(fileobj)
		log.Println("log demo")
		time.Sleep(time.Second * 3)
	}
}
