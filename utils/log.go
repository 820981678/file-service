package utils

import (
	"log"
	"os"
)

var LoadLogger *log.Logger

var ControllerLogger *log.Logger

func init() {
	logfile, err := os.OpenFile("systemlog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	//不能关闭，否则无法写入日志文件
	//defer logfile.Close()

	LoadLogger = log.New(logfile, "\r\n[load]", log.Ldate|log.Lmicroseconds)
	LoadLogger.Println("LoadLogger init success")

	ControllerLogger = log.New(logfile, "\r\n[controller]", log.Ldate|log.Lmicroseconds)
	LoadLogger.Println("ControllerLogger init success")
}