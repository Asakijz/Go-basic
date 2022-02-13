package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// get请求url？后面的是query参数
	// key=value格式，多个之间用&连接
	// ex：?query=小王子&age=18
	r.GET("/", func(c *gin.Context) {
		name := c.Query("query")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run(":9090")
}
