package database

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
	"github.com/borisbbtest/GoMon/internal/idm/database/postgres"
	pb "github.com/borisbbtest/GoMon/internal/idm/proto/idm"
)

type Storager interface {
	CreateTables(context.Context, *configs.AppConfig) error
	CreateSession(context.Context, *configs.AppConfig, *pb.Session) error
	DeleteSession(context.Context, *configs.AppConfig, string, string) error
	GetSession(context.Context, *configs.AppConfig, string, string) (*pb.Session, error)
	GetAllSessions(context.Context, *configs.AppConfig) ([]*pb.Session, error)
	CreateUser(context.Context, *configs.AppConfig, *pb.User) error
	DeleteUser(context.Context, *configs.AppConfig, string) error
	GetUser(context.Context, *configs.AppConfig, string) (*pb.User, error)
	GetAllUsers(context.Context, *configs.AppConfig) ([]*pb.User, error)
	Close()
}

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
