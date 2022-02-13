package main

import "github.com/gin-gonic/gin"

func sayhello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang",
	})
}

func main() {
	r := gin.Default()
	// 返回默认路由引擎
	r.GET("/", sayhello)
	// 指定用户使用get请求访问网页时，执行函数
	r.Run()
	// 启动服务
}
