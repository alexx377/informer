///////////////////////////////////////////////////////////
// Программа для сбора и передачи информации о системе
///////////////////////////////////////////////////////////
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Баннер при входе
func PrintBanner() {
	fmt.Println(STRING_BANNER)
}

//Точка входа
func main() {
	currentSettings := GetCommonSettings()

	PrintBanner()
	//Проверка доступности сервера
	serverOnLink := NetServerPing(currentSettings.Server)
	if serverOnLink {
		fmt.Println(STRING_SERVER_TEST, STRING_SERVER_AVAILABLE)
	} else {
		fmt.Println(STRING_SERVER_TEST, STRING_SERVER_NOT_AVAILABLE)
	}
	//----------------------------------------------
	fmt.Println(LINE)

	//Запрос расположения
	var location InfoLocation
	for location.Location == "" {
		//Ввод, пока не пройдет проверка на структуру и отсутствие дубликатов на сервере
		location = GetLocationInfo(currentSettings.Server, currentSettings.Prefixes, serverOnLink)
	}
	//----------------------------------------------
	fmt.Println(LINE)

	//Получение базовых данных о системе
	fmt.Println(STRING_STAGE_1)
	commonInfo := GetCommonInfo()

	//Получение сведений о сетевых адаптерах
	fmt.Println(STRING_STAGE_2)
	networkAdapterInfo := GetNetworkAdapterInfo()

	//Получение основных сведений о системе
	fmt.Println(STRING_STAGE_3)
	mainInfo := GetMainInfo()
	//----------------------------------------------
	fmt.Println(LINE)

	//Сведение данных в единую структуру
	mainInfo.Location = location.Location
	mainInfo.Description = location.Description

	mainInfo.Author = currentSettings.User

	mainInfo.ComputerName = commonInfo.ComputerName
	mainInfo.OS = commonInfo.OS
	mainInfo.Platform = commonInfo.Platform

	mainInfo.NetworkAdapters = networkAdapterInfo

	//Преобразование в строку JSON
	jsonInfo, err := json.Marshal(&mainInfo)
	if err != nil {
		fmt.Println(STRING_ERROR_JSON, err.Error())

		os.Exit(1)
	}

	//Заглушка
	fmt.Println(string(jsonInfo))
}
