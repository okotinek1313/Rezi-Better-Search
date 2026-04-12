package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port int `json:"port"`
}

func Create() {
	os.MkdirAll("./resources/shared", 0755)
	if _, err := os.Stat("./resources/shared/config.json"); os.IsNotExist(err) {
		file, err := os.Create("./resources/shared/config.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

func Read() Config {
	file, err := os.Open("./resources/shared/config.json")
	if err != nil {
		return Config{}
	}
	defer file.Close()
	var c Config
	json.NewDecoder(file).Decode(&c)
	return c
}

func AddConfig(c Config) {
	file, err := os.OpenFile("./resources/shared/config.json", os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	json.NewEncoder(file).Encode(c)
}
