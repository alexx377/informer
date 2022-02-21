///////////////////////////////////////////////////////////
// Программа для сбора и передачи информации о системе
///////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

//Баннер при входе
func PrintBanner() {
	fmt.Println(STRING_BANNER)
}

//Точка входа
func main() {
	var CurrentSettings CommonSettings
	CurrentSettings = GetCommonSettings()

	PrintBanner()
	//Проверка доступности сервера
	serverOnLink := NetServerPing(CurrentSettings.Server)
	if serverOnLink {
		fmt.Println(STRING_SERVER_TEST, STRING_SERVER_AVAILABLE)
	} else {
		fmt.Println(STRING_SERVER_TEST, STRING_SERVER_NOT_AVAILABLE)
	}
	//Запрос расположения
	var location InfoLocation
	for location.Location == "" {
		//Ввод, пока не пройдет проверка на структуру и отсутствие дубликатов на сервере
		location = GetLocationInfo(CurrentSettings.Server, serverOnLink)
	}

	fmt.Println(CurrentSettings.User)
	fmt.Println(location.Location)
}
