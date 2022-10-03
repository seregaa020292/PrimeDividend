package errmsg

const (
	AccessDenied               = "Доступ запрещен"
	UnknownError               = "Неизвестная ошибка"
	AuthorizationRequired      = "Требуется авторизация"
	ValidationError            = "При валидации возникла ошибка"
	ServerError                = "На сервере произошла ошибка"
	EncounteredRequestExternal = "При запросе на внешний ресурс возникла ошибка"
	MethodNotAllowed           = "Метод не разрешен"
	ValueMustNotEmpty          = "Значение не должно быть пустым"
	ValueAlreadyExists         = "Значение уже существует"
	CouldNotBeFound            = "Не удалось найти"
	FailedGetData              = "Не удалось получить данные"
	FailedAddData              = "Не удалось добавить данные"
	FailedUpdateData           = "Не удалось изменить данные"
	FailedDeleteData           = "Не удалось удалить данные"
	FailedSendMessage          = "Не удалось отправить сообщение"
	IsWaitStatus               = "В статусе ожидание"
	CheckTimeExpired           = "Время на проверку истекло"
	ConfirmWhileMatching       = "При сравнении возникла ошибка"
	CheckingWhileOccurred      = "При проверке произошла ошибка"
)
