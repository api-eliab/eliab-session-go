package main

import (
	"testing"

	_ "github.com/api-eliab/eliab-config-go"
)

func TestValidateCredentials(t *testing.T) {

	err := validateUser("eliezer.palacios@gmail.com")
	if err != nil {
		t.Error(err)
	}

	validate, err := validatePassword("email@email.com", "xxxxxxxxxx")
	if err != nil {
		t.Error(err)
		return
	}

	if !validate {
		t.Error("No se pudo validar!")
		return
	}

	t.Log("Te has logeado!")
}

func TestGetUserSons(t *testing.T) {

	sons, err := getUserSons(24)

	if err != nil {
		t.Error(err)
	}

	t.Error(sons)
}
