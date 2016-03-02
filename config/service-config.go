package config

import (
	"fmt"
	"encoding/json"
	"os"

	"file-service/utils"
)

type Config struct {
	ListenAddr string `json:"ListenAddr"`
	SaveRootPath string `json:"SaveRootPath"`
}

var ApplicationConfig Config

func init() {
	fmt.Println("loading config start...")
	utils.LoadLogger.Println("loading config start...")

	r, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	doc := json.NewDecoder(r)
	err = doc.Decode(&ApplicationConfig)
	if err != nil {
		panic(err)
	}

	fmt.Printf("config load success ApplicationConfig: %+v \n", ApplicationConfig)
	utils.LoadLogger.Println("config load success ApplicationConfig: ", ApplicationConfig)
}

