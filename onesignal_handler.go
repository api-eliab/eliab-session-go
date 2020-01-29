package main

import (
	"net/http"

	config "github.com/api-eliab/eliab-config-go"
	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

func sendBroadcastMessageHandler(w http.ResponseWriter, r *http.Request) {

	log.Info(config.Get.OneSignal.Key)

	messageRequest := MessageRequest{}

	request := apigo.Request{
		HTTPReq:    r,
		JSONStruct: &messageRequest,
	}

	response := request.UnmarshalBody()
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}

	response = sendMessageToUsersDevice(messageRequest.Message.Users, messageRequest.Message.Title, messageRequest.Message.Message, messageRequest.Message.Icon)

	apigo.SendResponse(response, w)
	return

}

func sendOneMessageHandler(w http.ResponseWriter, r *http.Request) {

	messageRequest := MessageRequest{}

	request := apigo.Request{
		HTTPReq:    r,
		JSONStruct: &messageRequest,
	}

	response := request.UnmarshalBody()
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}

	userID, response := request.GetURLParamInt64("userID")
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}

	messageRequest.Message.Users = append(messageRequest.Message.Users, userID)

	response = sendMessageToUsersDevice(messageRequest.Message.Users, messageRequest.Message.Title, messageRequest.Message.Message, messageRequest.Message.Icon)

	apigo.SendResponse(response, w)
	return

}
