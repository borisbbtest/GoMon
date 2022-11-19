package service

import "errors"

var (
	ErrEventWrongSeverity = errors.New("event have wrong severity, can't match") // Ошибка, информирующая о том, что у полученного события неизвестная критичность
	ErrEventWrongStatus   = errors.New("event have wrong status, can't match")   // Ошибка, информирующая о том, что у полученного события неизвестный статус
	ErrNoUserInContext    = errors.New("context don't have user information")    // Ошибка, информирующая о том, что контекст не содержит логин пользователя
)
