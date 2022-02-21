///////////////////////////////////////////////////////////
// Модуль строк выводимых сообщений
///////////////////////////////////////////////////////////

package main

const FILE_NAME_SETTINGS = "informer.json"
const DIRECTORY_NAME_FOR_SAVE = "data"

const STRING_BANNER = "=======================================================\n" +
	"| Утилита сбора краткой информации о системе. Вер.:1.00\n" +
	"======================================================="

const STRING_ERROR_READ_SETTINGS_FILE = "(W) Чтение файла настроек невозможно. Настройки установлены по-умолчанию."
const STRING_ERROR_SETTINGS_FILE = "(W) Ошибка обработки файла настроек. Настройки установлены по-умолчанию."
const STRING_ERROR_CONNECT_SERVER = "(W) Ошибка связи с сервером."
const STRING_ERROR_SERVER_RESPONCE = "(W) Неверный код ответа сервера."

const STRING_SERVER_TEST = "> Проверка связи с сервером:"
const STRING_SERVER_AVAILABLE = "Сервер доступен"
const STRING_SERVER_NOT_AVAILABLE = "Сервер недоступен"

const STRING_GET_ROOM = "Введите номер кабинета: "
const STRING_GET_ROOM_DESCR = "\t(Перед номером кабинета нужно указать код учебной площадки:\n" +
	"\tl68k - Ленина, 68\n" +
	"\tl68ak - Ленина, 68а\n" +
	"\ts27k - Союза Республик, 27\n" +
	"\ts22k - П.Сухова, 22\n" +
	"\ts71k - П.Сухова, 71\n" +
	"\ta84k - Э.Алексеевой, 84)"
const STRING_GET_WORKSPACE = "Введите номер рабочего места: "
const STRING_GET_WORKSPACE_DESCR = "\t(Перед номером кабинета нужно указать код принадлежности:\n" +
	"\ts - рабочее место студента\n" +
	"\tp - рабочее место преподавателя/сотрудника)"
const STRING_GET_DESCRIPTION = "Комментарии о рабочем месте: "
