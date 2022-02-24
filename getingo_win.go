///////////////////////////////////////////////////////////
// Модуль сбора информации о системе (Windows)
///////////////////////////////////////////////////////////

//go:build windows
// +build windows

package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

//Данные о материнской плате(WMI)
type Win32_BaseBoard struct {
	Manufacturer string
	Product      string
	SerialNumber string
}

//Данные о процессоре(WMI)
type Win32_Processor struct {
	Name          string
	MaxClockSpeed uint64
	NumberOfCores uint64
}

//Данные об оперативной памяти(WMI)
type Win32_PhysicalMemory struct {
	Manufacturer string
	SerialNumber string
	Capacity     uint64
}

//Данные о видеоконтроллере(WMI)
type Win32_VideoController struct {
	Name           string
	VideoProcessor string
	AdapterRAM     uint64
}

//Данные о жестком диске(WMI)
type Win32_DiskDrive struct {
	Model        string
	SerialNumber string
	Size         uint64
}

//Получение сведений о материнской плате
func GetBaseBoardInfo() []Win32_BaseBoard {
	var tmp []Win32_BaseBoard

	q := wmi.CreateQuery(&tmp, "")
	err := wmi.Query(q, &tmp)
	if err != nil {
		fmt.Println(STRING_BASE_BOARD_ERROR, err.Error())
	}

	return tmp
}

//Получение сведений о процессоре
func GetProcessorInfo() []Win32_Processor {
	var tmp []Win32_Processor

	q := wmi.CreateQuery(&tmp, "")
	err := wmi.Query(q, &tmp)
	if err != nil {
		fmt.Println(STRING_PROCESSOR_ERROR, err.Error())
	}

	return tmp
}

//Получение информации об оперативной памятью
func GetPhysicalMemorydInfo() []Win32_PhysicalMemory {
	var tmp []Win32_PhysicalMemory

	q := wmi.CreateQuery(&tmp, "")
	err := wmi.Query(q, &tmp)
	if err != nil {
		fmt.Println(STRING_MEMORY_ERROR, err.Error())
	}

	return tmp
}

//Получение данных о видеоконтроллерах
func GetVideoControllerInfo() []Win32_VideoController {
	var tmp []Win32_VideoController

	q := wmi.CreateQuery(&tmp, "")
	err := wmi.Query(q, &tmp)
	if err != nil {
		fmt.Println(STRING_VIDEO_ERROR, err.Error())
	}

	return tmp
}

//Получение данных о жестких дисках
func GetDiskDriveInfo() []Win32_DiskDrive {
	var tmp []Win32_DiskDrive

	q := wmi.CreateQuery(&tmp, "")
	err := wmi.Query(q, &tmp)
	if err != nil {
		fmt.Println(STRING_DISK_ERROR, err.Error())
	}

	return tmp
}

//Получение основных сведений о системе
func GetMainInfo() Info {
	var tmp Info

	tmp.ProcedureType = "Windows(WMI)"

	BaseBoard := GetBaseBoardInfo()
	Processor := GetProcessorInfo()
	PhysicalMemory := GetPhysicalMemorydInfo()
	VideoController := GetVideoControllerInfo()
	DiskDrive := GetDiskDriveInfo()

	//Обработка информации о материнской плате
	for _, v := range BaseBoard {
		tmp.BaseBoard.Model = v.Manufacturer + " " + v.Product
		tmp.BaseBoard.SerialNumber = v.SerialNumber
	}

	//Обработка информации о процессоре
	for _, v := range Processor {
		var Processor InfoProcessor

		Processor.Model = v.Name
		Processor.NumberOfCores = v.NumberOfCores
		Processor.Clock = v.MaxClockSpeed

		tmp.Processors = append(tmp.Processors, Processor)
	}

	//Обработка информации об оперативной памяти
	var TotalMemory uint64
	for _, v := range PhysicalMemory {
		var MemoryUnit InfoMemoryUnit

		MemoryUnit.Model = v.Manufacturer
		MemoryUnit.SerialNumber = v.SerialNumber
		MemoryUnit.Capacity = v.Capacity

		TotalMemory = TotalMemory + v.Capacity

		tmp.Memory.Units = append(tmp.Memory.Units, MemoryUnit)
	}
	tmp.Memory.TotalCapacity = TotalMemory

	//Обработка информации о видеоконтроллере
	for _, v := range VideoController {
		var VideoController InfoVideoController

		VideoController.Model = v.Name
		VideoController.Chip = v.VideoProcessor
		VideoController.Memory = v.AdapterRAM

		tmp.VideoControllers = append(tmp.VideoControllers, VideoController)
	}

	//Обработка информации о дисках
	for _, v := range DiskDrive {
		var HardDisk InfoHardDisk

		HardDisk.Model = v.Model
		HardDisk.SerialNumber = v.SerialNumber
		HardDisk.Capacity = v.Size

		tmp.HardDisks = append(tmp.HardDisks, HardDisk)
	}

	return tmp
}
