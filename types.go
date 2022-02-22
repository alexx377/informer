///////////////////////////////////////////////////////////
// Модуль описания типов данных
///////////////////////////////////////////////////////////

package main

//Префиксы мест расположения
type Prefix struct {
	Prefix      string `json:"prefix"`
	Description string `json:"description"`
}

//Общие настройки программы
type CommonSettings struct {
	Server   string   `json:"server"`
	Storage  string   `json:"storage"`
	User     string   `json:"user"`
	Prefixes []Prefix `json:"prefixes"`
}

//Данные о материнской плате
type InfoBaseBoard struct {
	Model        string `json:"model"`
	SerialNumber string `json:"serialNumber"`
}

//Данные о процессоре
type InfoProcessor struct {
	Model         string `json:"model"`
	NumberOfCores uint64 `json:"numberOfCores"`
	Clock         uint64 `json:"clock"`
}

//Данные об отдельном модуле оперативной памяти
type InfoMemoryUnit struct {
	Model        string `json:"model"`
	SerialNumber string `json:"serialNumber"`
	Capacity     uint64 `json:"capacity"`
}

//Данные об оперативной памяти вцелом
type InfoMemory struct {
	TotalCapacity uint64           `json:"totalCapacity"`
	Units         []InfoMemoryUnit `json:"units"`
}

//Данные о видеоконтроллере
type InfoVideoController struct {
	Model  string `json:"model"`
	Chip   string `json:"chip"`
	Memory uint64 `json:"memory"`
}

//Данные о жестком диске
type InfoHardDisk struct {
	Model        string `json:"model"`
	SerialNumber string `json:"serealNumber"`
	Capacity     uint64 `json:"capacity"`
}

//Данные о сетевом адаптере
type InfoNetworkAdapter struct {
	Name            string   `json:"name"`
	HardwareAddress string   `json:"hardwareAddress"`
	IpAddresses     []string `json:"ipAddresses"`
	IsUp            bool     `json:"isUp"`
}

//Итоговая информация о системе
type Info struct {
	Location         string                `json:"location"`
	Description      string                `json:"description"`
	Author           string                `json:"author"`
	ComputerName     string                `json:"computerName"`
	OS               string                `json:"os"`
	Platform         string                `json:"platform"`
	ProcedureType    string                `json:"procedureType"`
	BaseBoard        InfoBaseBoard         `json:"baseBoard"`
	Processors       []InfoProcessor       `json:"processors"`
	Memory           InfoMemory            `json:"memory"`
	VideoControllers []InfoVideoController `json:"videoControllers"`
	HardDisks        []InfoHardDisk        `json:"hardDisks"`
	NetworkAdapters  []InfoNetworkAdapter  `json:"networkAdapters"`
}

//Внутренние типы данных(не для экспотра)
//Имформация о расположении оборудования
type InfoLocation struct {
	Location    string
	Description string
}

//Общая информация о системе
type InfoCommon struct {
	ComputerName string
	OS           string
	Platform     string
}
