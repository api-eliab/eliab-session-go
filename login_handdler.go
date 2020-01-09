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

	if firstLogin == 1 {

		uuid := request.HTTPReq.Header.Get("DeviceUUID")
		appVersion := request.HTTPReq.Header.Get("AppVersion")
		osVersion := request.HTTPReq.Header.Get("OsVersion")
		os := request.HTTPReq.Header.Get("OS")
		deviceModel := request.HTTPReq.Header.Get("DeviceModel")
		timezoneStr := request.HTTPReq.Header.Get("Timezone")
		languaje := request.HTTPReq.Header.Get("Languaje")

		timezone, err := strconv.Atoi(timezoneStr)
		if err != nil {
			log.Error(err)
		}

		response = registerDeviceToPushNotification(uuid, appVersion, osVersion, os, deviceModel, languaje, timezone, loginRequest.Credentials.User)
		if response != nil {
			apigo.SendResponse(response, w)
			return
		}

	}

	response = validateCredentials(loginRequest.Credentials.User, loginRequest.Credentials.Password)
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}

	return
}
