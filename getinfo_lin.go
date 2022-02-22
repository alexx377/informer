///////////////////////////////////////////////////////////
// Модуль сбора информации о системе (Linux)
///////////////////////////////////////////////////////////

//go:build linux
// +build linux

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

//Получение сведений о материнской плате
func GetBaseBoardInfo() InfoBaseBoard {
	var tmp InfoBaseBoard

	//Производитель материнской платы
	baseBoardVendor, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/board_vendor")
	if err != nil {
		fmt.Println(STRING_BASE_BOARD_ERROR, err.Error())
	}

	//Модель материнской платы
	baseBoardModel, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/board_name")
	if err != nil {
		fmt.Println(STRING_BASE_BOARD_ERROR, err.Error())
	}

	//Серийный номер материнской платы
	baseBoardSerial, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/board_serial")
	if err != nil {
		fmt.Println(STRING_BASE_BOARD_ERROR, err.Error())
	}

	//Преобразование данных
	tmp.Model = strings.Trim(string(baseBoardVendor), " \r\n\t") + " " + strings.Trim(string(baseBoardModel), " \r\n\t")
	tmp.SerialNumber = strings.Trim(string(baseBoardSerial), " \r\n\t")

	return tmp
}

//Получение сведений о процессоре
func GetProcessorInfo() []InfoProcessor {
	var tmp []InfoProcessor

	//Процессор
	processor, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		fmt.Println(STRING_PROCESSOR_ERROR, err.Error())
	}

	//Получение данных о логических процессорах
	var logicalProcessorsParams []map[string]string

	logicalProcessorsInfo := strings.Split(strings.Trim(string(processor), " \r\n\t"), "\n\n")
	for _, logicalProcessorInfo := range logicalProcessorsInfo {
		logicalProcessorParams := strings.Split(logicalProcessorInfo, "\n")

		param := make(map[string]string)
		for _, paramString := range logicalProcessorParams {
			pair := strings.Split(paramString, ":")

			key := strings.Trim(pair[0], " \r\n\t")
			value := strings.Trim(pair[1], " \r\n\t")

			param[key] = value
		}

		logicalProcessorsParams = append(logicalProcessorsParams, param)
	}

	//Сведение логических процессоров к физическим
	id := int64(-1)
	for _, currentProcessor := range logicalProcessorsParams {
		cid, err := strconv.ParseInt(currentProcessor["core id"], 10, 64)
		if err != nil {
			fmt.Println(STRING_PROCESSOR_ERROR, err.Error())
		}

		//Если сменился id физического процессора, считать это новым физическим процессором
		if cid != id {
			var processorInfo InfoProcessor

			processorInfo.Model = currentProcessor["model name"]

			cores, err := strconv.ParseInt(currentProcessor["cpu cores"], 10, 64)
			if err != nil {
				fmt.Println(STRING_PROCESSOR_ERROR, err.Error())
			}
			processorInfo.NumberOfCores = uint64(cores)

			clock, err := strconv.ParseFloat(currentProcessor["cpu MHz"], 64)
			if err != nil {
				fmt.Println(STRING_PROCESSOR_ERROR, err.Error())
			}
			processorInfo.Clock = uint64(clock)

			id = cid

			tmp = append(tmp, processorInfo)
		}
	}

	return tmp
}

//Получение информации об оперативной памятью
func GetPhysicalMemorydInfo() InfoMemory {
	var tmp InfoMemory

	//Чтение информации об оперативной памяти
	memory, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println(STRING_MEMORY_ERROR, err.Error())
	}

	memoryInfoStrings := strings.Split(strings.Trim(string(memory), " \r\n\t"), "\n")
	memoryParam := make(map[string]string)

	for _, memoryInfoString := range memoryInfoStrings {
		pair := strings.Split(memoryInfoString, ":")

		key := strings.Trim(pair[0], " \r\n\t")
		value := strings.Trim(strings.ReplaceAll(pair[1], " kB", ""), " \r\n\t")

		memoryParam[key] = value
	}

	//Преобразование общего объема памяти
	totalMemory, err := strconv.ParseInt(memoryParam["MemTotal"], 10, 64)
	if err != nil {
		fmt.Println(STRING_MEMORY_ERROR, err.Error())
	}

	//Корректировка объема памяти(приведение к байтам)
	tmp.TotalCapacity = uint64(totalMemory) * 1024

	return tmp
}

//Получение данных о видеоконтроллерах
func GetVideoControllerInfo() []InfoVideoController {
	var tmp []InfoVideoController

	//Запрос перечисления всех видеоустройств в системе
	cmd := exec.Command("lspci")
	res, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(STRING_VIDEO_ERROR, err.Error())
	}
	defer res.Close()
	cmd.Start()

	//Чтение результатов запроса
	str, err := ioutil.ReadAll(res)
	if err != nil {
		fmt.Println(STRING_VIDEO_ERROR, err.Error())
	}

	//Обработка полученных данных
	videoInfoStrings := strings.Split(strings.Trim(string(str), " \r\n\t"), "\n")
	for _, videoInfoString := range videoInfoStrings {
		if strings.Index(videoInfoString, "VGA") < 0 {
			continue
		}

		var videoController InfoVideoController
		videoInfoString = videoInfoString[8:]
		pair := strings.Split(videoInfoString, ":")
		name := strings.Trim(pair[1], " /r/n/t")

		videoController.Model = name

		tmp = append(tmp, videoController)
	}

	return tmp
}

//Получение данных о жестких дисках
func GetDiskDriveInfo() []InfoHardDisk {
	var tmp []InfoHardDisk

	//Запрос перечисления всех блочных устройств в системе
	cmd := exec.Command("lsblk", "-lbdno", "SIZE,MODEL,SERIAL")
	res, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(STRING_DISK_ERROR, err.Error())
	}
	defer res.Close()
	cmd.Start()

	//Чтение результатов запроса
	str, err := ioutil.ReadAll(res)
	if err != nil {
		fmt.Println(STRING_DISK_ERROR, err.Error())
	}

	//Обработка полученных данных
	diskInfoStrings := strings.Split(strings.Trim(string(str), " \r\n\t"), "\n")
	for _, diskInfoString := range diskInfoStrings {
		var disk InfoHardDisk

		params := strings.Fields(diskInfoString)
		if len(params) < 3 {
			continue
		}

		disk.Model = params[1]
		disk.SerialNumber = params[2]
		size, err := strconv.ParseInt(params[0], 10, 64)
		if err != nil {
			fmt.Println(STRING_DISK_ERROR, err.Error())
		}
		disk.Capacity = uint64(size)

		tmp = append(tmp, disk)
	}

	return tmp
}

//Получение основных сведений о системе
func GetMainInfo() Info {
	var tmp Info

	tmp.ProcedureType = "Linux"

	tmp.BaseBoard = GetBaseBoardInfo()
	tmp.Processors = GetProcessorInfo()
	tmp.Memory = GetPhysicalMemorydInfo()
	tmp.VideoControllers = GetVideoControllerInfo()
	tmp.HardDisks = GetDiskDriveInfo()

	return tmp
}
