package main

import (
	"fmt"

	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

func validateCredentials(email, password string) apigo.Response {

	if err := validateUser(email); err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Credenciales incorrectas!",
			Message: "Usuario o contraseña incorectas",
		}
	}

	if !validatePassword(password) {
		return apigo.Error{
			Title:   "Credenciales incorrectas!",
			Message: "Usuario o contraseña incorrectas",
		}
	}

	user, err := getUserInfo(email)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información del usuario!",
			Message: "Error al consultar la información del usuario!",
		}
	}

	respData := ResponseLogin{}
	respData.User.ID = user.ID
	respData.User.FirstName = user.FirstName
	respData.User.LastName = fmt.Sprintf("%v %v", user.FirstLastName, user.SecondLastName)
	respData.User.Email = user.Email
	respData.User.Phone = user.Phone
	respData.User.Address = user.Address

	log.Info(user.ID)

	user.Sons, err = getUserSons(user.ID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de los hijos!",
			Message: "Error al consultar la información de los hijos!",
		}
	}

	for _, son := range user.Sons {
		var newSon ResponseSon
		newSon.ID = son.ID
		newSon.FirstName = son.FirstName
		newSon.LastName = fmt.Sprintf("%v %v", son.FirstLastName, son.SecondLastName)
		newSon.Avatar = son.Avatar
		respData.User.Sons = append(respData.User.Sons, newSon)
	}

	return apigo.Success{
		Content: respData,
	}
}

func validatePassword(password string) bool {
	return true
}
