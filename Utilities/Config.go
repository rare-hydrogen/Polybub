package Utilities

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Env           string `json:"env"`
	Connection    string `json:"connection"`
	CertPath      string `json:"certPath"`
	Pepper        string `json:"pepper"`
	Port          string `json:"port"`
	Domain        string `json:"domain"`
	TopDomain     string `json:"topDomain"`
	CookieName    string `json:"cookieName"`
	IsSecure      bool   `json:"isSecure"`
	ApiTitle      string `json:"apiTitle"`
	ApiVersion    string `json:"apiVersion"`
	SendgridName  string `json:"sendgridName"`
	SendgridEmail string `json:"sendgridEmail"`
	SendgridKey   string `json:"sendgridKey"`
}

var GlobalConfig Config

func GetConfig() Config {
	// Use the local file
	file, err := os.Open("config.json")
	if err != nil {
		// If no local overwrite, get from tmpfs
		file, err = os.Open("/var/www/app/Polybub/config.json")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Use the local file
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func GetBaseUrl(config Config) string {
	if config.Env == "production" {
		return "http://" + config.Domain + config.TopDomain
	} else {
		return "http://localhost" + ":" + config.Port
	}
}

func GetDomain(config Config) string {
	if config.Env == "production" {
		return config.Domain + config.TopDomain
	} else {
		return "localhost"
	}
}
