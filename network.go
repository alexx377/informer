///////////////////////////////////////////////////////////
// Модуль запроса информации о расположении оборудования
///////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"fmt"
	"io"
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

	if resp.StatusCode != http.StatusOK {
		fmt.Println(STRING_ERROR_SERVER_RESPONCE)

		return false
	}

	//Проверка тела ответа сервера
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(STRING_ERROR_SERVER_RESPONCE)

		return false
	}

	if string(respData) != "pong" {
		fmt.Println(STRING_ERROR_SERVER_RESPONCE)

		return false
	}

	//Все проверки пройдены
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

	if resp.StatusCode != http.StatusOK {
		fmt.Println(STRING_ERROR_SERVER_RESPONCE)

		return false
	}

	//Проверка тела ответа сервера
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(STRING_ERROR_SERVER_RESPONCE)

		return false
	}

	if string(respData) != "exist" {
		fmt.Println(STRING_ERROR_SERVER_RESPONCE)

		return false
	}
	return true
}
