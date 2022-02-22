///////////////////////////////////////////////////////////
// Модуль описания типов данных
///////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Функция генерации настроек по умолчанию.
func GetDefaultSettings() CommonSettings {
	var tmp CommonSettings

	tmp.Server = SERVER_ADDRESS
	tmp.Storage = DIRECTORY_NAME_FOR_SAVE
	tmp.User = USER_NAME

	return tmp
}

//Функция получения настроек из файла, или заполнение дефолтными настройками.
func GetCommonSettings() CommonSettings {
	var tmp CommonSettings

	//Проверка существования файла.
	_, err := os.Stat(FILE_NAME_SETTINGS)
	if err != nil {
		//Файл не найден или еще какая-то ошибка чтения файла. Настройки по-умолчанию.
		fmt.Println(STRING_ERROR_READ_SETTINGS_FILE, err.Error())

		return GetDefaultSettings()
	} else {
		//Ошибок нет. Попытка чтения.
		file, err := ioutil.ReadFile(FILE_NAME_SETTINGS)

		if err != nil {
			//Ошибка чтения файла. Настройки по-умолчанию.
			fmt.Println(STRING_ERROR_READ_SETTINGS_FILE, err.Error())

			return GetDefaultSettings()
		}

		err = json.Unmarshal(file, &tmp)
		if err != nil {
			//Ошибка структуры. Настройки по-умолчанию.
			fmt.Println(STRING_ERROR_SETTINGS_FILE, err.Error())

			return GetDefaultSettings()
		}

		//Настройки прочитаны.
		return tmp
	}
}
