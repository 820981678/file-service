package main

import (
	"fmt"
	"net/http"
	"file-service/controller"
	"file-service/utils"
)

func main() {
	
	utils.Init_config()
	
	utils.Init_log()
	utils.Init_domain()
	utils.Init_suffix()

	http.HandleFunc("/downfile", controller.DownFile)
	http.HandleFunc("/upfile", controller.UpFile)

	http.HandleFunc("/uptest", controller.UpTest)

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		fmt.Println("start error!")
	}
}
