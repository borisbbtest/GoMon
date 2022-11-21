package service

import "errors"

var (
	ErrMetricWrongType = errors.New("metric have wrong type, can't match") // Ошибка, информирующая о том, что у полученной метрики неизвестный тип
	ErrNoUserInContext = errors.New("context don't have user information") // Ошибка, информирующая о том, что контекст не содержит логин пользователя
)
