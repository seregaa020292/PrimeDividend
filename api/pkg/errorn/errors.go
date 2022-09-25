package errorn

import "net/http"

const (
	TargetUnknown Target = 1 + iota
	TargetValidate
	TargetDb
	TargetParse
	TargetEmail
	TargetAuth
	TargetFound
	TargetServer
	TargetInvalid
)

var (
	ErrorAuthorization = NewError(TargetAuth, http.StatusBadRequest, "Ошибка авторизации")
	ErrorAccess        = NewError(TargetAuth, http.StatusForbidden, "Ошибка доступа")
	ErrorUnAuthorized  = NewError(TargetAuth, http.StatusUnauthorized, "Ошибка требуется авторизация")

	ErrorNotFoundElement = NewError(TargetFound, http.StatusNotFound, "Не удалось найти элемент")

	ErrorSelect = NewError(TargetDb, http.StatusBadRequest, "Ошибка получения данных")
	ErrorInsert = NewError(TargetDb, http.StatusBadRequest, "Ошибка добавления данных")
	ErrorUpdate = NewError(TargetDb, http.StatusBadRequest, "Ошибка обновления данных")
	ErrorDelete = NewError(TargetDb, http.StatusBadRequest, "Ошибка удаления данных")

	ErrorSendEmail  = NewError(TargetEmail, http.StatusBadRequest, "Ошибка отправки email сообщения")
	ErrorExistEmail = NewError(TargetEmail, http.StatusBadRequest, "Такой email уже существует")

	ErrorUserNoConfirm = NewError(TargetInvalid, http.StatusBadRequest, "Пользователь не подтвержден")

	ErrorValidate = NewError(TargetValidate, http.StatusBadRequest, "Ошибка валидации")

	ErrorUnknown = NewError(TargetUnknown, http.StatusInternalServerError, "Неизвестная ошибка")

	ErrorPasswordIncorrect = NewError(TargetInvalid, http.StatusBadRequest, "Пароль не верный")
)
