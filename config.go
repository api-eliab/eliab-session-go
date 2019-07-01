package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

var config Configuration

type Configuration struct {
	Mode     string
	General  General
	DataBase DataBase
}

type General struct {
	ServerAddress string
}

type DataBase struct {
	User     string
	Password string
	Server   string
	DataBase string
	Port     int64
}

func loadConfiguration() {
	path := "./config/config.toml"

	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Printf("Couldn't read config file at [%s]\n", path)
		log.Fatal(err)
	}
}
