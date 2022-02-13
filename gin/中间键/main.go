package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//定义一个中间键m1统计请求处理函数耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	start := time.Now()
	// c.Next() //调用后续的处理函数
	c.Abort()//阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", m1, index)
	r.Run()
}
