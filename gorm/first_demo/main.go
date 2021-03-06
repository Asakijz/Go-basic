package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Userinfo struct {
	Id     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//自动迁移
	db.AutoMigrate(&Userinfo{})
	// u1 := Userinfo{1, "小王子", "男", "篮球"}
	// db.Create(&u1)
	//查询
	var u Userinfo
	db.First(&u)
	fmt.Printf("%#v\n", u)
	//更新
	db.Model(&u).Update("Hobby", "蛙泳")

}
