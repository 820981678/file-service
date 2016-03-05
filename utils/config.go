package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	ListenAddr   string `json:"ListenAddr"`
	SaveRootPath string `json:"SaveRootPath"`
}

var ApplicationConfig Config

func Init_config() {
	LoadLogger.Infoln("loading config start...")

	r, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	doc := json.NewDecoder(r)
	err = doc.Decode(&ApplicationConfig)
	if err != nil {
		panic(err)
	}

	LoadLogger.Infof("config load success ApplicationConfig: %+v \n", ApplicationConfig)
}
