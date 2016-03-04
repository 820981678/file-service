package main

import (
	"fmt"
	"net/http"
	"os"

	"file-service/controller"
	"file-service/global"

	_"file-service/config"
)

func checkSaveDir() {
	fmt.Println(global.ApplicationConfig.SaveRootPath)
	info, err := os.Stat(global.ApplicationConfig.SaveRootPath)

	temp := true
	if err != nil {
		temp = os.IsExist(err)
	} else {
		temp = info.IsDir()
	}

	if !temp {
		err := os.Mkdir(global.ApplicationConfig.SaveRootPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		fmt.Println("create dir success")
	} else {
		fmt.Println("have dir")
	}
}

func main() {
	checkSaveDir()

	fmt.Println("file service start...")
	http.HandleFunc("/downfile", controller.DownFile)
	http.HandleFunc("/upfile", controller.UpFile)

	http.HandleFunc("/uptest", controller.UpTest)

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		fmt.Println("start error!")
	}
}