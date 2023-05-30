package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type config struct {
	Repo string `json:"repo"`
	Port string `json:"port"`
}

func NewConfig() config {
	var path string
	flag.StringVar(&path, "path", "\\config.json", "Path to configuration")
	flag.Parse()
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	cfg := config{}
	err := decoder.Decode(&cfg)
	if err != nil {
		log.Fatalln("Ошибка чтения конфигов", err)
	}
	return cfg
}
