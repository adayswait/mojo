package mlog

import (
	"log"
	"os"
)

var mlog *log.Logger

var Log = log.Println

func init() {
	mojolog, _ := os.OpenFile("./mojo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //打开日志文件，不存在则创建
	mlog = log.New(mojolog, "", log.Ldate|log.Ltime|log.Lshortfile)
	Log = mlog.Println
}
