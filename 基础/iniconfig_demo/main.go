package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type mysqlconfig struct {
	address  string `ini:"address"`
	port     int    `ini:"port"`
	username string `ini:"username"`
	password string `ini:"password"`
}

type redisconfig struct {
	host     string `ini:"host"`
	port     int    `ini:"port"`
	password string `ini:"password"`
	database int    `ini:"database"`
}

type config struct {
	mysqlconfig `ini:"mysql"`
	redisconfig `ini:"redis"`
}

func loadini(filename string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	//判断传进来的值是否为指针类型
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer")
		return
	}
	//判断传进来的值是否为结构体指针类型
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a struct")
		return
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	//读取文件到变量暂存并且按换行符分割
	lineslice := strings.Split(string(b), "\r\n")
	// fmt.Printf("%#v\n", lineslice)
	var structname string
	for index, line := range lineslice {
		//去掉字符串的空格
		line = strings.TrimSpace(line)
		//跳过无任何字段的空行
		if len(line) == 0 {
			continue
		}
		//如果读到注释则跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//如果是[开头则表示是section
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return
			}
			sectionname := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionname) == 0 {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return
			}
			//根据字符串去data里面根据反射来找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionname == field.Tag.Get("ini") {
					structname = field.Name
					fmt.Printf("找到了%s对应的嵌套结构体%s\n", sectionname, structname)
				}
			}
		} else {
			//如果该字段行以=开头或者不包含=则报错
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line%d syntax error", index+1)
				return
			}
			//以=分割，取出该行对应的key和value
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			// value := strings.TrimSpace(line[index+1:])
			//判断传进来的data是否为结构体
			v := reflect.ValueOf(data)
			svalue := v.Elem().FieldByName(structname)
			stype := svalue.Type()
			// structobj := v.Elem().FieldByName(structname)
			if stype.Kind() != reflect.Struct {
				err = fmt.Errorf("data name:%dshould be a struct", structname)
				return
			}
			var fieldname string
			var filetype reflect.StructField
			for i := 0; i < svalue.NumField(); i++ {
				filed := stype.Field(i)
				filetype = filed
				if filed.Tag.Get("ini") == key {
					fieldname = filed.Name
					break
				}
			}
			// fileobj := svalue.FieldByName(fieldname)
			fmt.Println(fieldname, filetype.Type.Kind())
		}

	}
	return
}

func main() {
	var mc config
	err := loadini("./config.ini", &mc)
	if err != nil {
		fmt.Printf("open config file failed,err:%v", err)
		return
	}
	fmt.Println(mc)
}
