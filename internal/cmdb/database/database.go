// Package database описывает интерфейс работы с БД и создает объект репозитория в соответствии с настройками
package database

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/cmdb/configs"
	"github.com/borisbbtest/GoMon/internal/cmdb/database/postgres"
	pb "github.com/borisbbtest/GoMon/internal/models/cmdb"
)

// Storager - интерфейс, описывающий работу с хранилищем idm.
type Storager interface {
	CreateTables(context.Context, *configs.AppConfig) error                // создает таблицы при старте
	CreateObject(context.Context, *configs.AppConfig, *pb.Ci) error        // создание новой КЕ
	GetObject(context.Context, *configs.AppConfig, string) (*pb.Ci, error) // получение существующей КЕ
	DeleteObject(context.Context, *configs.AppConfig, string) error        // удаление существующей КЕ

	Close() // остановка работы с репозиторием
}

// NewDBStorage - создает хранилище на основе параметров сервера.
func NewDBStorage(ctx context.Context, cfg *configs.AppConfig) (Storager, error) {
	if true { // В будущем можно будет добавить другие поддерживаемые СУБД
		repo, err := postgres.NewCmdbRepo(ctx, cfg)
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
