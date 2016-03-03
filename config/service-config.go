package config

import (
	"encoding/json"
	"os"

	"file-service/utils"
	"file-service/global"
)

func init() {
	utils.LoadLogger.Infoln("loading config start...")

	r, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	doc := json.NewDecoder(r)
	err = doc.Decode(&global.ApplicationConfig)
	if err != nil {
		panic(err)
	}

	utils.LoadLogger.Infof("config load success ApplicationConfig: %+v \n", global.ApplicationConfig)
}

