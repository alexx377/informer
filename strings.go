///////////////////////////////////////////////////////////
// Модуль строк выводимых сообщений
///////////////////////////////////////////////////////////

package main

const FILE_NAME_SETTINGS = "informer.json"
const SERVER_ADDRESS = "127.0.0.1:3030"
const DIRECTORY_NAME_FOR_SAVE = "data"
const USER_NAME = "Unknown"

const URL_PING = "/ping"
const URL_CHECK_LOCATION = "/check"
const URL_SEND_DATA = "/send"

const NAME_DELIMITER = "-"
const LINE = "----------"

const STRING_BANNER = "=======================================================\n" +
	"| Утилита сбора краткой информации о системе. Вер.:1.00\n" +
	"======================================================="

const STRING_ERROR_READ_SETTINGS_FILE = "(ES) Чтение файла настроек невозможно. Настройки установлены по-умолчанию."
const STRING_ERROR_SETTINGS_FILE = "(ES) Ошибка обработки файла настроек. Настройки установлены по-умолчанию."
const STRING_ERROR_CONNECT_SERVER = "(ES) Ошибка связи с сервером."
const STRING_ERROR_SERVER_RESPONCE = "(ES) Неверный код ответа сервера."

const STRING_ERROR_LOCATION_STRUCT = "(EL) Ошибка в структуре описания расположения."
const STRING_ERROR_LOCATION_EXIST = "(EL) Информация о данном расположении уже имеется на сервере."

const STRING_ERROR_COMMON_INFO = "(EcI) Ошибка получения общих сведений о системе."
const STRING_ERROR_NETWORK_ADAPTERS = "(EN) Ошибка перечисления сетевых адаптеров."
const STRING_ERROR_NETWORK_ADDRESSES = "(EN) Ошибка получения IP-адресов интерфейса."

const STRING_BASE_BOARD_ERROR = "(EmI) Ошибка получение данных о материнской плате."
const STRING_PROCESSOR_ERROR = "(EmI) Ошибка получение данных о процессоре."
const STRING_MEMORY_ERROR = "(EmM) Ошибка получение данных об оперативной памяти."
const STRING_VIDEO_ERROR = "(EmV) Ошибка получение данных о видеоконтроллерах."
const STRING_DISK_ERROR = "(EmH) Ошибка получение данных о жестких дисках."

const STRING_LOCAL_SAVE_ERROR = "(EF) Ошибка локального сохранения."

const STRING_ERROR_JSON = "(E) Ошибка преобразования итоговых данных."

const STRING_SERVER_TEST = "> Проверка связи с сервером:"
const STRING_SERVER_AVAILABLE = "Сервер доступен"
const STRING_SERVER_NOT_AVAILABLE = "Сервер недоступен"

const STRING_GET_LOCATION_CODE_DESCR = "Доступные коды корпусов:"
const STRING_GET_LOCATION_CODE = "Введите код корпуса: "
const STRING_GET_ROOM = "Введите номер кабинета: "
const STRING_GET_WORKSPACE_CODE_DESCR = "Доступные коды рабочих мест:\n" +
	"\ts - рабочее место студента\n" +
	"\tp - рабочее место преподавателя/сотрудника"
const STRING_GET_WORKSPACE_CODE = "Введите код рабочего места: "
const STRING_GET_WORKSPACE = "Введите номер рабочего места: "
const STRING_GET_DESCRIPTION = "Введите комментарий к рабочему месту: "
const STRING_LOCATION_SUMMARY = "Код расположения: "
const STRING_LOCATION_DESCRIPTION_SUMMARY = "Комментарий к рабочему месту:"
const STRING_LOCATION_NO_COMMENT = "[Без комментариев]"

const STRING_STAGE_1 = "Шаг 1: Сбор базовых данных о системе..."
const STRING_STAGE_2 = "Шаг 2: Сбор данных о сетевых адаптерах..."
const STRING_STAGE_3 = "Шаг 3: Сбор данных об аппаратной части..."

const STRING_DATA_LOCAL_SAVED = "(Ok) Данные сохранены локально."
const STRING_DATA_REMOTE_SAVED = "(Ok) Данные переданы на сервер."

const STRING_READY_MESSAGE = "Данные собраны успешно."
