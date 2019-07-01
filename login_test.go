package main

import (
	"testing"
)

func TestValidateCredentials(t *testing.T) {
	loadConfiguration()
	dbConnect()

	err := validateUser("eliezer.palacios@gmail.com")
	if err != nil {
		t.Error(err)
	}
	if !validatePassword("xxxxxxxxxx") {
		t.Error("No se pudo validar!")
	}

	t.Log("Te has logeado!")
}

func TestGetUserSons(t *testing.T) {
	loadConfiguration()
	dbConnect()

	sons, err := getUserSons(24)

	if err != nil {
		t.Error(err)
	}

	t.Error(sons)
}
