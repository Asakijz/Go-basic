package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//定义模型
type User struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	//连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//把模型和数据库中的表对应起来
	db.AutoMigrate(&User{})
	//创建
	// u1 := User{Name: "小王子", Age: 18}
	// db.Create(&u1)
	// u2 := User{Name: "七米", Age: 20}
	// db.Create(&u2)
	//查询
	// var user User
	// db.First(&user) //first是以ID为关键字查询，没有则无法查询(可接第二个参数，仅当id为整数才行db.first(&user,10))
	// fmt.Printf("%v\n", user)

	// var users []User
	// db.Find(&users) //find查询全部结果，需传入切片类型变量
	// fmt.Printf("%v\n", users)

	//更新
	// user.Name = "沙河"
	// user.Age = 30
	// db.Save(&user)                        //默认修改所有字段
	// db.Model(&user).Update("name", "小王子") //只修改对应字段 update只修改一个字段

	// m1 := map[string]interface{}{
	// 	"Name": "周杰伦",
	// 	"Age":  16,
	// }
	// db.Model(&user).Select("Name").Updates(m1) //选择一个字段更新
	// db.Model(&user).Omit("Name").Updates(m1)   //排除一个字段更新

	//让表格中所有用户的年龄在原来的基础上+2

	// db.Model(&user).Update("Age", gorm.Expr("Age+?", 2))

	//删除-->软删除 保留字段但查询会跳过

	var u = User{}
	u.ID = 1
	db.Delete(&u)
	db.Where("Name=?", "七米").Delete(&User{}) //根据指定条件删除
}
