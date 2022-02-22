///////////////////////////////////////////////////////////
// Программа для сбора и передачи информации о системе
///////////////////////////////////////////////////////////
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

//Баннер при входе
func PrintBanner() {
	fmt.Println(STRING_BANNER)
}

//Сохранение данных в файл(резервное)
func SaveFile(Folder string, FileName string, Data string) {
	//Проверка существования директории
	fileInfo, err := os.Stat(Folder)
	if err != nil {
		if os.IsNotExist(err) {
			//Директория не существует, нужно создать
			err = os.Mkdir(Folder, 0777)
			if err != nil {
				//Директория не создана
				fmt.Println(STRING_LOCAL_SAVE_ERROR, err.Error())
				return
			}
		} else {
			//Другие ошибки
			fmt.Println(STRING_LOCAL_SAVE_ERROR, err.Error())
			return
		}
	}

	//Проверка, точно ли по указанному пути директория
	if fileInfo.IsDir() {
		//Директория существует, пробуем сохранять файл
		file, err := os.Create(path.Join(Folder, FileName))
		if err != nil {
			fmt.Println(STRING_LOCAL_SAVE_ERROR, err.Error())
			return
		}
		defer file.Close()

		_, err = file.WriteString(string(Data))
		if err != nil {
			fmt.Println(STRING_LOCAL_SAVE_ERROR, err.Error())
			return
		}
	}
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

	//Сохранение собранных данных в локальный файл
	//SaveFile(currentSettings.Storage, mainInfo.Location+".json", string(jsonInfo))
	fmt.Println(STRING_DATA_LOCAL_SAVED)

	//При наличии связи, передача данных на сервер
	if serverOnLink {
		SendDataToServer(currentSettings.Server, jsonInfo)
	}

	//Завершение
	fmt.Println(LINE)
	fmt.Println(STRING_READY_MESSAGE)
}
