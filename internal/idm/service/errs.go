package service

import "errors"

var (
	ErrUserExists     = errors.New("user already exist") // Ошибка, информирующая о том, что создаваемый пользователь уже существует
	ErrEmptySQLResult = errors.New("no result in sql")   // Ошибка, информирующая о том, что запрашиваемый элемент отсутствует
	ErrWrongPassword  = errors.New("WrongPassword")      // Ошибка, информирующая о том, что представлена неверная пара login/password
)
