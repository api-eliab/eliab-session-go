package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	config "github.com/api-eliab/eliab-config-go"
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

	validation, err := validatePassword(email, password)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al validar las credenciales",
			Message: "Error al autenticar el usuario",
		}
	}

	if !validation {
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
		newSon.Section = son.Section
		newSon.Grade = son.Grade
		respData.User.Sons = append(respData.User.Sons, newSon)
	}

	return apigo.Success{
		Content: respData,
	}
}

func validatePassword(email, password string) (bool, error) {

	url := fmt.Sprintf("%v?key=%v&email=%v&password=%v", config.Get.Services["Authentication"].URL, config.Get.AuthenticationToken, email, password)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return false, err
	}

	var c = &http.Client{}

	response, err := c.Do(req)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return false, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	log.Info(string(data))

	if string(data) == "Invalid Credentials" {
		return false, nil
	}

	if string(data) == "Invalid Token" {
		return false, nil
	}

	userID, err := strconv.Atoi(string(data))
	if err != nil {
		return false, err
	}

	if userID == 0 {
		return false, errors.New("Ha ocurrido un error inesperado")
	}

	return true, nil

}

func validateAppVersion(inAppVersionStr string) apigo.Response {

	log.Debug("Device AppVersion: ", inAppVersionStr)
	log.Debug("Latest AppVersion: ", latestSupportedAppVersion)

	inAppVersion, err := convertVersionThreePartStrToInt(inAppVersionStr)

	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title: "No se ha podido validar la versión instalada en tu dispositivo",
		}
	}

	if inAppVersion.compareLatestVersion(latestSupportedAppVersion) {
		return nil
	}

	return apigo.Error{
		Title:   "¡Nueva actualización! Ingresa al Playstore y actualiza para ingresar",
		Message: "Debes actualizar tu aplicación para poder continuar",
		Action:  "https://play.google.com/store/apps/details?id=school.palacios.gt.com.schoolapp&hl=en",
	}

}

// VersionGNUStandar doc ...
type VersionGNUStandar struct {
	Major     int
	Minor     int
	Reversion int
}

func (currentVersion VersionGNUStandar) compareLatestVersion(latestVerstion VersionGNUStandar) bool {

	if currentVersion.Major > latestVerstion.Major {

		return true

	} else if currentVersion.Major == latestVerstion.Major {

		if currentVersion.Minor > latestVerstion.Minor {

			return true

		} else if currentVersion.Minor == latestVerstion.Minor {

			if currentVersion.Reversion >= latestVerstion.Reversion {

				return true

			}

		}

	}

	return false

}

var latestSupportedAppVersion VersionGNUStandar

func convertLatestSupportedAppVersionThreePart() error {

	var err error

	latestSupportedAppVersion, err = convertVersionThreePartStrToInt(config.Get.AppVersion)

	if err != nil {
		return err
	}

	return nil

}

func convertVersionThreePartStrToInt(versionStr string) (version VersionGNUStandar, err error) {

	versionThreePart := strings.Split(versionStr, ".")

	version.Major, err = strconv.Atoi(versionThreePart[0])
	version.Minor, err = strconv.Atoi(versionThreePart[1])
	version.Reversion, err = strconv.Atoi(versionThreePart[2])

	return

}
