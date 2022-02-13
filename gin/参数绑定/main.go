package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Userinfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		// username := c.Query("username") //根据？后面的键值对来返回值
		// password := c.Query("password")
		// u := userinfo{
		// 	username: username,
		// 	password: password,
		// }
		// fmt.Printf("%#v\n", u)
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "ok",
		// })
		var u Userinfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
		fmt.Printf("%#v\n", u)
	})
	r.Run(":9090")
}
