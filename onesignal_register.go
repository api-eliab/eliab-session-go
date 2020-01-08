package main

import (
	config "github.com/api-eliab/eliab-config-go"
	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/jonesignal"
	"github.com/josuegiron/log"
)

func registerDeviceToPushNotification(uuid string, userID int64) apigo.Response {

	deviceReq := jonesignal.OSAddDeviceReq{
		AppID:      config.Get.OneSignal.AppID,
		Identifier: uuid,
	}

	osID, err := jonesignal.AddDevice(deviceReq)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "No se pudo agregar el dispositivo al proveedor!",
			Message: "No se pudo agregar al dispositivo al proveedor!",
		}
	}

	log.Info(osID)

	return nil

}
