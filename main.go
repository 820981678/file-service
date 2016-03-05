package main

import (
	"fmt"
	"net/http"
	"os"

	"file-service/controller"

	"file-service/utils"
)

func checkSaveDir() {

	info, err := os.Stat(utils.ApplicationConfig.SaveRootPath)

	temp := true
	if err != nil {
		temp = os.IsExist(err)
	} else {
		temp = info.IsDir()
	}

	if !temp {
		err := os.Mkdir(utils.ApplicationConfig.SaveRootPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		fmt.Println("create dir " + utils.ApplicationConfig.SaveRootPath)
	} else {
		fmt.Println("have dir " + utils.ApplicationConfig.SaveRootPath)
	}
}

func main() {
	utils.Init_log()
	utils.Init_suffix()
	utils.Init_config()

	checkSaveDir()

	http.HandleFunc("/downfile", controller.DownFile)
	http.HandleFunc("/upfile", controller.UpFile)

	http.HandleFunc("/uptest", controller.UpTest)

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		fmt.Println("start error!")
	}
}
