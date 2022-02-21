///////////////////////////////////////////////////////////
// Модуль запроса информации о расположении оборудования
///////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"fmt"
	"net/http"
)

//Проверка доступности сервера
func NetServerPing(ServerAddr string) bool {
	resp, err := http.Get("http://" + ServerAddr + "/ping")
	if err != nil {
		fmt.Println(STRING_ERROR_CONNECT_SERVER, err.Error())

		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(STRING_ERROR_SERVER_RESPONCE)

		return false
	}

	//TODO:Реализовать проверку тела ответа
	return true
}

//Проверка существования расположения
func NetLocationExist(ServerAddr string, Location string) bool {
	req := []byte(`{ "location": "` + Location + `" }`)

	resp, err := http.Post("http://"+ServerAddr+"/check", "", bytes.NewBuffer(req))
	if err != nil {
		fmt.Println(STRING_ERROR_CONNECT_SERVER, err.Error())

		return false
	}

	if resp.StatusCode != 200 {
		fmt.Println(STRING_ERROR_SERVER_RESPONCE)

		return false
	}

	//TODO:Реализовать проверку тела ответа
	return true
}
