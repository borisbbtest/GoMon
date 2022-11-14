package service

import "errors"

var (
	ErrObjectExists   = errors.New("object already exist")             // Ошибка, информирующая о том, что создаваемый объект уже существует
	ErrEmptySQLResult = errors.New("no result in sql")                 // Ошибка, информирующая о том, что запрашиваемый элемент отсутствует
	ErrWrongPassword  = errors.New("WrongPassword")                    // Ошибка, информирующая о том, что представлена неверная пара login/password
	ErrInsertObjects  = errors.New("one or more Cis was not created")  // Ошибка, информирующая о том, что есть КЕ в Batch, которые не создались в БД
	ErrSelectObjects  = errors.New("one or more Cis was not selected") // Ошибка, информирующая о том, что есть КЕ в Batch, которые не были получены из БД, выборка не полная
	ErrDeleteObjects  = errors.New("one or more Cis was not deleted")  // Ошибка, информирующая о том, что есть КЕ в Batch, которые не были удалены из БД
)
