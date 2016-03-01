package main

import (
	"fmt"
	"net/http"

	"os"
	"io"
	"strings"

	"controller"
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

var savefilename int32 = 0

func main() {
	fmt.Println("file service start...")
	http.HandleFunc("/downfile", controller.downFile)
	http.HandleFunc("/upfile", controller.upFile)

	http.HandleFunc("/uptest", controller.upTest)

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		fmt.Println("start error!")
	}
}