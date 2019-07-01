package main

import (
	"testing"

	"github.com/josuegiron/log"
)

func TestLoadConfig(t *testing.T) {
	loadConfiguration()
	log.Println(config.DataBase.User)
}
