// Package database описывает интерфейс работы с БД и создает объект репозитория в соответствии с настройками
package database

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
	"github.com/borisbbtest/GoMon/internal/idm/database/postgres"
	pb "github.com/borisbbtest/GoMon/internal/models/idm"
)

// Storager - интерфейс, описывающий работу с хранилищем idm.
type Storager interface {
	CreateTables(context.Context, *configs.AppConfig) error                              // создает таблицы при старте
	CreateSession(context.Context, *configs.AppConfig, *pb.Session) error                // создание новой сессиии
	DeleteSession(context.Context, *configs.AppConfig, string, string) error             // удаление сессии
	GetSession(context.Context, *configs.AppConfig, string, string) (*pb.Session, error) // получение существующей сессии
	GetAllSessions(context.Context, *configs.AppConfig) ([]*pb.Session, error)           // получение списка всех сессий
	CreateUser(context.Context, *configs.AppConfig, *pb.User) error                      // создание нового пользователя
	DeleteUser(context.Context, *configs.AppConfig, string) error                        // удаление пользователя
	GetUser(context.Context, *configs.AppConfig, string) (*pb.User, error)               // получение существующего пользователя
	GetAllUsers(context.Context, *configs.AppConfig) ([]*pb.User, error)                 // получение списка всех сессий
	Close()                                                                              // остановка работы с репозиторием
}

// NewDBStorage - создает хранилище на основе параметров сервера.
func NewDBStorage(ctx context.Context, cfg *configs.AppConfig) (Storager, error) {
	if true { // В будущем можно будет добавить другие поддерживаемые СУБД
		repo, err := postgres.NewIdmRepo(ctx, cfg)
		if err != nil {
			return nil, err
		}
		err = repo.CreateTables(ctx, cfg)
		if err != nil {
			return nil, err
		}
		return repo, nil
	}
	return nil, nil
}
