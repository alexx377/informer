///////////////////////////////////////////////////////////
// Модуль запроса информации о расположении оборудования
///////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Проверка структуры расположения
func LocationStructTest(Location string) bool {
	//TODO: Реализовать проверку по шаблону
	return true
}

//Замена функции scanf (работат одинаково в Windows и Linux)
func GetString() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

//Получение информации о расположении от пользователя
func GetLocationInfo(ServerAddr string, ServerExist bool) InfoLocation {
	var tmp InfoLocation

	fmt.Println(STRING_GET_ROOM)
	fmt.Println(STRING_GET_ROOM_DESCR)
	fmt.Print(":")
	room := GetString()
	fmt.Println(STRING_GET_WORKSPACE)
	fmt.Println(STRING_GET_WORKSPACE_DESCR)
	fmt.Print(":")
	workspace := GetString()
	tmp.Location = strings.ToLower(room + "-" + workspace)

	if LocationStructTest(tmp.Location) == false {
		fmt.Println(STRING_ERROR_LOCATION_STRUCT)

		tmp.Location = ""
		tmp.Description = ""

		return tmp
	}

	if ServerExist {
		if NetLocationExist(ServerAddr, tmp.Location) {
			fmt.Println(STRING_ERROR_LOCATION_EXIST)

			tmp.Location = ""
			tmp.Description = ""

			return tmp
		}
	}

	fmt.Print(STRING_GET_DESCRIPTION)
	description := GetString()
	tmp.Description = description

	return tmp
}
