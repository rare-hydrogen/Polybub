package Utilities

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Env        string `json:"env"`
	Connection string `json:"connection"`
	Pepper     string `json:"pepper"`
	Port       string `json:"port"`
	Hostname   string `json:"hostname"`
	ApiTitle   string `json:"apiTitle"`
	ApiVersion string `json:"apiVersion"`
}

var GlobalConfig Config

func GetConfig() Config {
	return GetConfigByPath("config.json")
}

func GetConfigByPath(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func GetCurrentEnv(config Config) string {
	if config.Env == "production" {
		return "http://" + config.Hostname
	} else {
		return "http://localhost" + ":" + config.Port
	}
}
