package main

import (
	"encoding/json"
	"net/http"

	apigolang "github.com/josuegiron/api-golang"
)

func login(w http.ResponseWriter, r *http.Request) {

	request := apigolang.Request{
		HTTPReq:    r,
		JSONStruct: LoginRequest{},
	}

	response := request.UnmarshalBody()
	if response != nil {
		apigolang.SendResponse(response, w)
	}

	muckData := []byte(`
	{
		"user":{
			"id":1821, 
			"firstName": "Firulais",
			"lastName": "Canelo",
			"email": "firulais@colegios.edu",
			"phone": "12345678",
			"address": "12av. 19-33, Zona 12",
			"sons": [
				{
					"id": 1892,
					"firstName": "Manchas",
					"lastName": "Canelo",
					"avatar": 3
				},
				{
					"id": 1893,
					"firstName": "Perla",
					"lastName": "Canelo",
					"avatar": 4
				}
			]
		}
	}`)

	muckStruct := LoginResponse{}

	if err := json.Unmarshal(muckData, &muckStruct); err != nil {
		panic(err)
	}

	// Set Sesson
	w.Header().Set("SessionId", "MySession")

	apigolang.SuccesContentResponse("Inicio de sesión", "¡Esta es información de prueba!", muckStruct, w)
}
