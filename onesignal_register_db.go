package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jgolang/mysqltools"
)

func saveUserDeviceInDB(uuid, appVersion, osVersion, os, deviceModel, language string, timezone int, userID string, playerID string) error {

	query := fmt.Sprintf("INSERT INTO user_device (uuid, appVersion, osVersion, os, deviceModel, timezoneStr, languaje, timezone, onesignal_id, status, user_id, created_at, updated_at) VALUES ( @uuid, @appVersion, @osVersion, @os, @deviceModel, @timezone, @languaje, @timezone, @onesignalID, @status, (select id from mas_person where email = @userID), NOW(), NOW())")
	query2, err := mysqltools.GetQueryString(
		query,
		sql.Named("uuid", uuid),
		sql.Named("appVersion", appVersion),
		sql.Named("osVersion", osVersion),
		sql.Named("os", os),
		sql.Named("deviceModel", deviceModel),
		sql.Named("timezone", timezone),
		sql.Named("languaje", language),
		sql.Named("onesignalID", playerID),
		sql.Named("status", "active"),
		sql.Named("userID", userID),
	)

	if err != nil {
		return err
	}

	result, err := db.Exec(query2)
	if err != nil {
		return err
	}

	numRowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if numRowAffected == 0 {
		return errors.New("No se ingresó la tupla")
	}

	return nil

}

func getUserDevicesOneSignal(userID int64) (onesignalIDs []string, err error) {
	var query2 string
	query := fmt.Sprintf("SELECT onesignal_id FROM user_device WHERE user_id = @userID")
	query2, err = mysqltools.GetQueryString(
		query,
		sql.Named("userID", userID),
	)
	if err != nil {
		return
	}

	log.Println(query2)

	rows, errR := db.Query(query2)
	if errR != nil {
		return onesignalIDs, errR
	}

	for rows.Next() {
		var onesignalID string
		err = rows.Scan(
			&onesignalID,
		)
		if err != nil {
			return
		}

		onesignalIDs = append(onesignalIDs, onesignalID)
	}

	return
}
