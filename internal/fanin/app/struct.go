// Package app собирает приложение. Считывает конфиг, поднимает http сервер. Останавливает по получению сигнала от ОС
package app

import "google.golang.org/grpc"

// ConnPool - пул подключений gRPC
type ConnPool struct {
	idm     *grpc.ClientConn // подключение к модулю idm (управление пользователями)
	cmdb    *grpc.ClientConn // подключение к модулю cmdb (управление КЕ)
	metrics *grpc.ClientConn // подключение к модулю metrics (хранение метрик)
	events  *grpc.ClientConn // подключение к модулю events (хранение событий)
}
