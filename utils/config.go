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
