package utils

import (
	"encoding/json"
	"os"
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
		panic(err)
	}

	doc := json.NewDecoder(r)
	err = doc.Decode(&ApplicationConfig)
	if err != nil {
		panic(err)
	}

}
