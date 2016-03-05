package utils

import (
	"fmt"
	"log"
	"os"
)

var LoadLogger *MyLogger

var ControllerLogger *MyLogger

func Init_log() {
	logfile, err := os.OpenFile("systemlog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	//不能关闭，否则无法写入日志文件
	//defer logfile.Close()

	l := log.New(logfile, "\r\n[load]", log.Ldate|log.Lmicroseconds)
	LoadLogger = &MyLogger{l}
	LoadLogger.Infoln("LoadLogger init success")

	c := log.New(logfile, "\r\n[controller]", log.Ldate|log.Lmicroseconds)
	ControllerLogger = &MyLogger{c}
	LoadLogger.Infoln("ControllerLogger init success")
}

type MyLogger struct {
	*log.Logger
}

func (c *MyLogger) Infoln(v interface{}) {
	c.Println(v)
	fmt.Println(v)
}

func (c *MyLogger) Infof(s string, v interface{}) {
	c.Printf(s, v)
	fmt.Printf(s, v)
}
