package errorn

import "net/http"

var (
	ErrorAuthorization = NewError(TargetAuth, http.StatusBadRequest, "Ошибка авторизации")

	ErrorSelect = NewError(TargetDb, http.StatusBadRequest, "Ошибка получения данных")
	ErrorInsert = NewError(TargetDb, http.StatusBadRequest, "Ошибка добавления данных")
	ErrorUpdate = NewError(TargetDb, http.StatusBadRequest, "Ошибка обновления данных")
	ErrorDelete = NewError(TargetDb, http.StatusBadRequest, "Ошибка удаления данных")

	ErrorValidate = NewError(TargetValidate, http.StatusBadRequest, "Ошибка валидации")

	ErrorUnknown = NewError(TargetUnknown, http.StatusInternalServerError, "Неизвестная ошибка")
)
