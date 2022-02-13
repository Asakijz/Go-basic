package mylogger

import (
	"fmt"
	"time"
)

//往终端输出日志

type consolelogger struct {
	level loglever
}

//获取级别并返回
func Newlog(levelstr string) consolelogger {
	lever, err := parseloglevel(levelstr)
	if err != nil {
		panic(err)
	}
	return consolelogger{
		level: lever,
	}
}

//开关判断是否输出结果
func (c consolelogger) enable(loglever loglever) bool {
	return c.level <= loglever
}

func (c consolelogger) clog(lv loglever, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcname, filename, lineno := getinfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getlogstring(lv), filename, funcname, lineno, msg)
	}
}

func (c consolelogger) Debug(format string, a ...interface{}) {
	c.clog(DEBUG, format, a...)
}

func (c consolelogger) Trace(format string, a ...interface{}) {
	c.clog(TRACE, format, a...)
}

func (c consolelogger) Info(format string, a ...interface{}) {
	c.clog(INFO, format, a...)
}
func (c consolelogger) Warning(format string, a ...interface{}) {
	c.clog(WARNING, format, a...)
}
func (c consolelogger) Error(format string, a ...interface{}) {
	c.clog(ERROR, format, a...)
}
func (c consolelogger) Fatal(format string, a ...interface{}) {
	c.clog(FATAL, format, a...)
}
