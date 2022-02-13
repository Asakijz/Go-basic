package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// var wg sync.WaitGroup

//往文件中写入日志
type filelogger struct {
	level       loglever
	filepath    string //日志文件路径
	filename    string //日志文件名称
	fileobj     *os.File
	errfileobj  *os.File
	maxfilesize int64
}

func Newfilelogger(levelstr, fp, fn string, maxsize int64) *filelogger {
	loglever, err := parseloglevel(levelstr)
	if err != nil {
		panic(err)
	}
	f1 := &filelogger{
		level:       loglever,
		filepath:    fp,
		filename:    fn,
		maxfilesize: maxsize,
	}
	err = f1.initfile() //按照文件路径何名字打开
	if err != nil {
		panic(err)
	}
	return f1
}

func (f *filelogger) initfile() error {
	//文件名称拼接
	fullfilename := path.Join(f.filepath, f.filename)
	fileobj, err := os.OpenFile(fullfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	errfileobj, err := os.OpenFile(fullfilename+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	f.fileobj = fileobj
	f.errfileobj = errfileobj
	return nil

}

func (f *filelogger) enable(loglever loglever) bool {
	return f.level <= loglever
}

func (f *filelogger) checksize(file *os.File) bool {
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	//当前文件大于等于定义的文件最大值则返回true
	return fileinfo.Size() >= f.maxfilesize
}

func (f *filelogger) splitfile(file *os.File) (*os.File, error) {
	//切割日志文件
	nowstr := time.Now().Format("20060102")
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get fileinfo failed,err:%v\n", err)
		return nil, err
	}
	logname := path.Join(f.filepath, fileinfo.Name())         //旧的日志文件路径
	newlogname := fmt.Sprintf("%s.backup%s", logname, nowstr) //拼接后新的日志文件名称
	//1、关闭当前日志文件
	file.Close()
	//2、备份并rename  xx.log->xx.log.backup20211123
	os.Rename(logname, newlogname)
	//3、打开新的日志文件
	fileobj, err := os.OpenFile(logname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil, err
	}
	//4、将新的日志文件赋值给f.fileobj
	return fileobj, nil
}

func (f *filelogger) flog(lv loglever, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcname, filename, lineno := getinfo(3)
		if f.checksize(f.fileobj) {
			newfile, err := f.splitfile(f.fileobj)
			if err != nil {
				return
			}
			f.fileobj = newfile
		}
		fmt.Fprintf(f.fileobj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getlogstring(lv), filename, funcname, lineno, msg)
		//如果日志等级大于ERROR则在err日志文件中再记录一遍
		if lv >= ERROR {
			if f.checksize(f.errfileobj) {
				newfile, err := f.splitfile(f.errfileobj)
				if err != nil {
					return
				}
				f.errfileobj = newfile
			}
		}
		fmt.Fprintf(f.errfileobj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getlogstring(lv), filename, funcname, lineno, msg)
		// defer wg.Done()
	}
}

func (f *filelogger) Debug(format string, a ...interface{}) {
	f.flog(DEBUG, format, a...)
}

func (f *filelogger) Trace(format string, a ...interface{}) {
	f.flog(TRACE, format, a...)
}

func (f *filelogger) Info(format string, a ...interface{}) {
	f.flog(INFO, format, a...)
}
func (f *filelogger) Warning(format string, a ...interface{}) {
	f.flog(WARNING, format, a...)
}
func (f *filelogger) Error(format string, a ...interface{}) {
	f.flog(ERROR, format, a...)
}
func (f *filelogger) Fatal(format string, a ...interface{}) {
	f.flog(FATAL, format, a...)
}

// func (f *filelogger) close() {
// 	f.fileobj.Close()
// 	f.errfileobj.Close()
// }
