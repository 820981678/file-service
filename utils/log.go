package utils

import (
	"fmt"
	"log"
	"os"
)

var LoadLogger *MyLogger

var ControllerLogger *MyLogger

func Init_log() {
	fmt.Println("load log start......")
	
	logfile, err := os.OpenFile(ApplicationConfig.LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	//不能关闭，否则无法写入日志文件
	//defer logfile.Close()

	l := log.New(logfile, "\r\n[load]", log.Ldate|log.Lmicroseconds)
	LoadLogger = &MyLogger{l}
	fmt.Println("load LoadLogger complete")

	c := log.New(logfile, "\r\n[controller]", log.Ldate|log.Lmicroseconds)
	ControllerLogger = &MyLogger{c}
	fmt.Println("load ControllerLogger complete")
}

type MyLogger struct {
	*log.Logger
}

func (c *MyLogger) Infoln(v interface{}) {
	c.Println(v)
	fmt.Println(v)
}

func (c *MyLogger) Infof(s string, v ...interface{}) {
	c.Printf(s, v)
	fmt.Printf(s, v)
}
