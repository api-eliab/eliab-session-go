package main

import (
	config "github.com/api-eliab/eliab-config-go"
	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/jonesignal"
	"github.com/josuegiron/log"
)

func registerDeviceToPushNotification(uuid, appVersion, osVersion, os, deviceModel, language string, timezone int, userID string) apigo.Response {

	deviceReq := jonesignal.OSAddDeviceReq{
		AppID:       config.Get.OneSignal.AppID,
		Identifier:  uuid,
		DeviceModel: deviceModel,
		DeviceOs:    osVersion,
		Timezone:    timezone,
		Language:    language,
		GameVersion: appVersion,
		DeviceType:  1, // Android
	}

	onesignalID, err := jonesignal.AddDevice(deviceReq)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "No se pudo agregar el dispositivo al proveedor!",
			Message: "No se pudo agregar al dispositivo al proveedor!",
		}
	}

	err = saveUserDeviceInDB(uuid, appVersion, osVersion, os, deviceModel, language, timezone, userID, onesignalID)
	if err != nil {
		log.Error(err)
		// return apigo.Error{
		// 	Title:   "No se pudo asociar el dispositivo!",
		// 	Message: "No se pudo asocial el dispositivo!",
		// }
	}

	return nil

}

func sendMessageToUsersDevice(users []int64, title, message, iconURL string) apigo.Response {

	notification := jonesignal.NotificationRequest{
		AppID: config.Get.OneSignal.AppID,
		Headings: jonesignal.Headings{
			En: title,
			Es: title,
		},
		Contents: jonesignal.Contents{
			En: message,
			Es: message,
		},
		LargeIcon: iconURL,
	}

	for _, userID := range users {

		ids, err := getUserDevicesOneSignal(userID)
		if err != nil {
			log.Error(err)
			return apigo.Error{
				Title:   "No fue posible enviar el mensaje!",
				Message: "No fue posible enviar el mensaje!",
			}
		}

		notification.IncludePlayerIds = append(notification.IncludePlayerIds, ids...)

	}

	log.Info(notification)

	id, err := jonesignal.SendNotification(notification)
	if err != nil {
		return apigo.Error{
			Title:   "No fue posible enviar el mensaje!",
			Message: "No fue posible enviar el mensaje!",
		}
	}

	log.Info(id)

	return apigo.Success{
		Title:   "Mensaje enviado correctamente!",
		Message: "Mensaje enviado correctamente!",
	}

}
