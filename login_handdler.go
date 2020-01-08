package main

import (
	"net/http"

	apigo "github.com/josuegiron/api-golang"
)

func login(w http.ResponseWriter, r *http.Request) {

	loginRequest := LoginRequest{}

	request := apigo.Request{
		HTTPReq:    r,
		JSONStruct: &loginRequest,
	}

	firstLogin, response := request.GetQueryParamInt64("firstLogin")
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}

	if firstLogin == 1 {

		uuid := request.HTTPReq.Header.Get("UUID")

		response = registerDeviceToPushNotification(uuid, request.UserID)
		if response != nil {
			apigo.SendResponse(response, w)
			return
		}

	}

	response = request.UnmarshalBody()
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}

	response = loginRequest.ValidationFields()
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}

	response = validateCredentials(loginRequest.Credentials.User, loginRequest.Credentials.Password)
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}

	return
}
