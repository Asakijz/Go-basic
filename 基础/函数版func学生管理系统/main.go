package main

import (
	"fmt"
	"os"
)

//函数版学生管理系统，实现查看、增加、删除学生功能
var (
	allstudent map[int64]*student
)

//构造体函数
func newstudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

type student struct {
	id   int64
	name string
}

func showallstudent() {
	for k, v := range allstudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

func addstudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)
	newstu := newstudent(id, name)
	allstudent[id] = newstu
}

func deletestudent() {
	var deleteid int64
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&deleteid)
	delete(allstudent, deleteid)
}

func main() {
	allstudent = make(map[int64]*student, 50) //开辟内存空间
	for {
		//打印菜单
		fmt.Println("欢迎来到学生管理系统")
		fmt.Println(`
		1、查看所有学生
		2、新增学生
		3、删除学生
		4、退出菜单
	`)
		//获取用户选项
		fmt.Print("请输入你要选择的选项：")
		var choice int
		fmt.Scanln(&choice)
		// fmt.Printf("你选择了%d这个选项", choice)
		//执行对应的函数
		switch choice {
		case 1:
			showallstudent()
		case 2:
			addstudent()
		case 3:
			deletestudent()
		case 4:
			os.Exit(1) //os.Exit(0)程序正常运行，反正出错
		default:
			fmt.Println("请输入正确的菜单选项")
		}

	}
}
