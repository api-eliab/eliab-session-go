package main

import (
	"regexp"

	apigolang "github.com/josuegiron/api-golang"
)

// LoginRequest is struct json request
type LoginRequest struct {
	Credentials struct {
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"credentials"`
}

// ValidationFields validate fields
func (login *LoginRequest) ValidationFields() apigolang.Response {

	if login.Credentials.User == "" {
		return apigolang.Error{
			Title:   "Campo usuario vacío",
			Message: "Debes ingresar tu usuario",
		}
	}

	if login.Credentials.Password == "" {
		return apigolang.Error{
			Title:   "Campo password vacío",
			Message: "Debes ingresar una contraseña",
		}
	}

	if !validateEmail(login.Credentials.User) {
		return apigolang.Error{
			Title:   "Formato de usuario inválido",
			Message: "Ingresa un correo electrónico valido",
		}
	}

	return nil
}

//	Validate Email
func validateEmail(email string) bool {
	//	Regular expression
	re := regexp.MustCompile(`^[K-Za-z0-9._%+\-]+@[K-Za-z0-9.\-]+\.[a-z]{1,4}$`)
	return re.MatchString(email)
}
