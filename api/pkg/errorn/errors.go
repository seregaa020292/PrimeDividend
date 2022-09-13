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
)

var (
	ErrorAuthorization = NewError(TargetAuth, http.StatusBadRequest, "Ошибка авторизации")
	ErrorAccess        = NewError(TargetAuth, http.StatusForbidden, "Ошибка доступа")

	ErrorFoundElement = NewError(TargetFound, http.StatusNotFound, "Не удалось найти элемент")

	ErrorSelect = NewError(TargetDb, http.StatusBadRequest, "Ошибка получения данных")
	ErrorInsert = NewError(TargetDb, http.StatusBadRequest, "Ошибка добавления данных")
	ErrorUpdate = NewError(TargetDb, http.StatusBadRequest, "Ошибка обновления данных")
	ErrorDelete = NewError(TargetDb, http.StatusBadRequest, "Ошибка удаления данных")

	ErrorValidate = NewError(TargetValidate, http.StatusBadRequest, "Ошибка валидации")

	ErrorUnknown = NewError(TargetUnknown, http.StatusInternalServerError, "Неизвестная ошибка")
)
