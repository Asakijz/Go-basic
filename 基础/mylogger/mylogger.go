package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//自定义日志库

type loglever uint16

//定义日志级别
const (
	UNKNOW loglever = iota //iota初始为0，后续逐渐累加
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//将级别转换小写
func parseloglevel(s string) (loglever, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	}
	err := errors.New("错误的日志级别")
	return UNKNOW, err
}

//获取代码信息行号
func getinfo(skip int) (funcname, filename string, lineno int) {
	pc, file, lineno, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime caller() failed")
		return

	}
	funcname = runtime.FuncForPC(pc).Name()
	filename = path.Base(file)
	funcname = strings.Split(funcname, ".")[1]
	return
}

//获取级别并返回字符串
func getlogstring(lv loglever) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}
