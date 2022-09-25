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
	ErrForbidden    = NewError(TargetAuth, http.StatusForbidden, "Доступ запрещен")
	ErrUnauthorized = NewError(TargetAuth, http.StatusUnauthorized, "Требуется авторизация")

	ErrMethodNotAllowed = NewError(TargetFound, http.StatusMethodNotAllowed, "Метод не разрешен")
	ErrNotFound         = NewError(TargetFound, http.StatusNotFound, "Не удалось найти")

	ErrSelect = NewError(TargetDb, http.StatusBadRequest, "Ошибка получения данных")
	ErrInsert = NewError(TargetDb, http.StatusBadRequest, "Ошибка добавления данных")
	ErrUpdate = NewError(TargetDb, http.StatusBadRequest, "Ошибка обновления данных")
	ErrDelete = NewError(TargetDb, http.StatusBadRequest, "Ошибка удаления данных")

	ErrSendEmail  = NewError(TargetEmail, http.StatusBadRequest, "Ошибка отправки email сообщения")
	ErrExistEmail = NewError(TargetEmail, http.StatusBadRequest, "Такой email уже существует")

	ErrUserNoConfirm = NewError(TargetInvalid, http.StatusBadRequest, "Пользователь не подтвержден")

	ErrValidate = NewError(TargetValidate, http.StatusBadRequest, "Ошибка валидации")

	ErrUnknown = NewError(TargetUnknown, http.StatusInternalServerError, "Неизвестная ошибка")

	ErrPasswordIncorrect = NewError(TargetInvalid, http.StatusBadRequest, "Пароль не верный")
)
