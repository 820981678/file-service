package utils

import (
	"encoding/json"
	"os"
	"fmt"
	"runtime"
)

type Config struct {
	ListenAddr   string `json:"ListenAddr"`
	SaveRootPath string `json:"SaveRootPath"`
	LogFilePath  string `json:"LogFilePath"`
}

var ApplicationConfig Config

func Init_config() {
	fmt.Println("load config start......")
	
	r, err := os.Open("config.json")
	if err != nil {
		fmt.Println("file config.json error use default config")
		def()
		return
	}

	doc := json.NewDecoder(r)
	err = doc.Decode(&ApplicationConfig)
	if err != nil {
		fmt.Println("file config.json error use default config")
		def()
	}
	
	checkSaveDir()
	fmt.Printf("load config complete ApplicationConfig: %+v \n", ApplicationConfig)
}

func checkSaveDir() {

	info, err := os.Stat(ApplicationConfig.SaveRootPath)

	temp := true
	if err != nil {
		temp = os.IsExist(err)
	} else {
		temp = info.IsDir()
	}

	if !temp {
		err := os.Mkdir(ApplicationConfig.SaveRootPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		fmt.Println("create dir " + ApplicationConfig.SaveRootPath)
	} else {
		fmt.Println("have dir " + ApplicationConfig.SaveRootPath)
	}
}

func def() {
	ApplicationConfig = Config {
			ListenAddr : "80",
	}
	sys := runtime.GOOS
	switch sys {
		case "linux":
			ApplicationConfig.SaveRootPath = "/home/file-service-data/"
			ApplicationConfig.LogFilePath = "/home/file-service-data/log.log"
		case "windows":
			ApplicationConfig.SaveRootPath = "c://file-service-data/"
			ApplicationConfig.LogFilePath = "c://file-service-data/log.log"
		default:
			ApplicationConfig.SaveRootPath = "/home/file-service-data/"
			ApplicationConfig.LogFilePath = "/home/file-service-data/log.log"
	}
}
