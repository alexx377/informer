///////////////////////////////////////////////////////////
// Модуль запроса информации о расположении оборудования
///////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Замена функции scanf (работат одинаково в Windows и Linux)
func GetString() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

//Получение информации о расположении от пользователя
func GetLocationInfo(ServerAddr string, PrefixList []Prefix, ServerExist bool) InfoLocation {
	var tmp InfoLocation

	//Ввод кода расположения
	fmt.Println(STRING_GET_LOCATION_CODE_DESCR)
	for i, currentPrefix := range PrefixList {
		fmt.Println("\t[" + strconv.Itoa(i) + "] " + currentPrefix.Prefix + " - " + currentPrefix.Description)
	}
	locationCode := -1
	for locationCode < 0 {
		fmt.Print(STRING_GET_LOCATION_CODE)
		tmp, err := strconv.Atoi(GetString())
		if err == nil {
			if (tmp >= 0) && (tmp < len(PrefixList)) {
				locationCode = tmp
			}
		}
	}

	//Ввод номера кабинета
	room := -1
	for room < 0 {
		fmt.Print(STRING_GET_ROOM)
		tmp, err := strconv.Atoi(GetString())
		if err == nil {
			if tmp >= 0 {
				room = tmp
			}
		}
	}

	//Ввод типа рабочего места
	fmt.Println(STRING_GET_WORKSPACE_CODE_DESCR)
	var workspaceCode string
	for workspaceCode == "" {
		fmt.Print(STRING_GET_WORKSPACE_CODE)
		tmp := strings.ToLower(GetString())

		if (tmp == "s") || (tmp == "p") {
			workspaceCode = tmp
		}
	}

	//Ввод номера рабочего места
	workspace := -1
	for workspace < 0 {
		fmt.Print(STRING_GET_WORKSPACE)
		tmp, err := strconv.Atoi(GetString())
		if err == nil {
			if tmp >= 0 {
				workspace = tmp
			}
		}
	}

	//Формируем полное имя расположения
	tmp.Location = PrefixList[locationCode].Prefix + NAME_DELIMITER + strconv.Itoa(room) + NAME_DELIMITER + workspaceCode + NAME_DELIMITER + strconv.Itoa(workspace)

	//При наличии связи с сервером, проверка дубликатов расположений
	if ServerExist {
		if NetLocationExist(ServerAddr, tmp.Location) {
			fmt.Println(STRING_ERROR_LOCATION_EXIST)

			tmp.Location = ""
			tmp.Description = ""

			return tmp
		}
	}

	//Ввод комментария
	fmt.Print(STRING_GET_DESCRIPTION)
	description := GetString()
	tmp.Description = description

	//Отображение итогового имени
	fmt.Println(LINE)
	fmt.Println(STRING_LOCATION_SUMMARY, tmp.Location)
	fmt.Println(STRING_LOCATION_DESCRIPTION_SUMMARY)
	if tmp.Description != "" {
		fmt.Println(tmp.Description)
	} else {
		fmt.Println(STRING_LOCATION_NO_COMMENT)
	}

	return tmp
}
