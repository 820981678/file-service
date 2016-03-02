package main

import (
	"fmt"
	"net/http"
	"os"

	"file-service/controller"
	_ "file-service/config"
)

func init() {
	info, err := os.Stat("D://fileservice")

	temp := true
	if err != nil {
		temp = os.IsExist(err)
	} else {
		temp = info.IsDir()
	}

	if !temp {
		err := os.Mkdir("D://fileservice", os.ModePerm)
		if err != nil {
			panic(err)
		}
		fmt.Println("create dir success")
	} else {
		fmt.Println("have dir")
	}
}

func main() {
	fmt.Println("file service start...")
	http.HandleFunc("/downfile", controller.DownFile)
	http.HandleFunc("/upfile", controller.UpFile)

	http.HandleFunc("/uptest", controller.UpTest)

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		fmt.Println("start error!")
	}
}