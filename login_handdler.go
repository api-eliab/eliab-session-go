package main

import (
	"net/http"
	"strconv"

	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

func login(w http.ResponseWriter, r *http.Request) {

	loginRequest := LoginRequest{}

	request := apigo.Request{
		HTTPReq:    r,
		JSONStruct: &loginRequest,
	}

	response := request.UnmarshalBody()
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}

	response = loginRequest.ValidationFields()
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}

	firstLogin, response := request.GetQueryParamInt64("firstLogin")
	appVersionHeader := request.HTTPReq.Header.Get("AppVersion")

	if appVersionHeader != appVersion {

		log.Debug(appVersion)

		response := apigo.Error{
			Title:   "¡Nueva actualización! Ingresa al Playstore y actualiza para ingresar",
			Message: "Debes actualizar tu aplicación para poder continuar",
			Action:  "https://play.google.com/store/apps/details?id=school.palacios.gt.com.schoolapp&hl=en",
		}

		apigo.SendResponse(response, w)

		return

	}

	if firstLogin == 1 {

		uuid := request.HTTPReq.Header.Get("DeviceUUID")
		osVersion := request.HTTPReq.Header.Get("OsVersion")
		os := request.HTTPReq.Header.Get("OS")
		deviceModel := request.HTTPReq.Header.Get("DeviceModel")
		timezoneStr := request.HTTPReq.Header.Get("Timezone")
		languaje := request.HTTPReq.Header.Get("Languaje")
		playerID := request.HTTPReq.Header.Get("PlayerID")

		timezone, err := strconv.Atoi(timezoneStr)
		if err != nil {
			log.Error(err)
		}

		registerDeviceToPushNotification(uuid, appVersion, osVersion, os, deviceModel, languaje, timezone, loginRequest.Credentials.User, playerID)

	}

	response = validateCredentials(loginRequest.Credentials.User, loginRequest.Credentials.Password)
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}

	return
}
