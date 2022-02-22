///////////////////////////////////////////////////////////
// Модуль получения базовой информации
///////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"net"

	"github.com/matishsiao/goInfo"
)

//Получение базовых сведений о системе
func GetCommonInfo() InfoCommon {
	var tmp InfoCommon

	osInfo, err := goInfo.GetInfo()
	if err != nil {
		fmt.Println(STRING_ERROR_COMMON_INFO, err.Error())

		return tmp
	}

	//Преобразование базовых сведений о системе к нужному формату
	tmp.ComputerName = osInfo.Hostname
	tmp.OS = osInfo.Kernel + " " + osInfo.Core
	tmp.Platform = osInfo.Platform

	return tmp
}

//Получение сведений о сетевых адаптерах
func GetNetworkAdapterInfo() []InfoNetworkAdapter {
	var tmp []InfoNetworkAdapter

	//Получение списка сетевых адаптеров
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(STRING_ERROR_NETWORK_ADAPTERS, err.Error())

		return tmp
	}

	//Перебор адаптеров, получение подробных сведений о каждом
	for _, iface := range ifaces {
		var currentNetworkAdapter InfoNetworkAdapter

		//Если интерфейс loopback - игнорируем
		if (iface.Flags & net.FlagLoopback) != 0 {
			continue
		}

		//Преобразование сведений о адаптере
		currentNetworkAdapter.Name = iface.Name
		currentNetworkAdapter.HardwareAddress = iface.HardwareAddr.String()
		if (iface.Flags & net.FlagUp) != 0 {
			currentNetworkAdapter.IsUp = true
		} else {
			currentNetworkAdapter.IsUp = false
		}

		//Получение IP-адресов текущего адаптера
		addresses, err := iface.Addrs()
		if err == nil {
			for _, address := range addresses {
				currentNetworkAdapter.IpAddresses = append(currentNetworkAdapter.IpAddresses, address.String())
			}
		} else {
			fmt.Println(STRING_ERROR_NETWORK_ADDRESSES, err.Error())
		}

		tmp = append(tmp, currentNetworkAdapter)
	}

	return tmp
}
