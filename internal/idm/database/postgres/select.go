package postgres

import (
	"context"
	"embed"
	"errors"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
	"github.com/borisbbtest/GoMon/internal/idm/service"
	pb "github.com/borisbbtest/GoMon/internal/models/idm"
	"github.com/jackc/pgx/v4"
)

// Файлы SQL для вставки записей в таблицы хранятся в директории migrations/select/
//
//go:embed migrations/select/*.sql
var SQLSelect embed.FS

// GetUser - функция, которая возвращает пользователя в структуре protobuf c указанным Login
func (r *IdmRepo) GetUser(ctx context.Context, cfg *configs.AppConfig, login string) (*pb.User, error) {
	sqlBytes, err := SQLSelect.ReadFile("migrations/select/SQLSelectUser.sql")
	if err != nil {
		return nil, err
	}
	sqlQuery := string(sqlBytes)
	row := r.Pool.QueryRow(ctx, sqlQuery, login)
	var user PGUser
	err = row.Scan(&user.Id, &user.Login, &user.Firstname, &user.Lastname, &user.Password, &user.Source, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, service.ErrEmptySQLResult
		}
		log.Error().Err(err).Msg("failed get user")
		return nil, err
	}
	result := user.ConvertToPB()
	return result, nil
}

// GetAllUsers - функция, которая возвращает всех пользователей в структуре protobuf
func (r *IdmRepo) GetAllUsers(ctx context.Context, cfg *configs.AppConfig) ([]*pb.User, error) {
	var list []*pb.User
	sqlBytes, err := SQLSelect.ReadFile("migrations/select/SQLSelectAllUsers.sql")
	if err != nil {
		return nil, err
	}
	sqlQuery := string(sqlBytes)
	rows, err := r.Pool.Query(ctx, sqlQuery)
	if err != nil {
		log.Error().Err(err).Msg("execute query failed")
		return list, err
	}
	defer rows.Close()
	for rows.Next() {
		var user PGUser
		err = rows.Scan(&user.Id, &user.Login, &user.Firstname, &user.Lastname, &user.Password, &user.Source, &user.CreatedAt)
		if err != nil {
			log.Error().Err(err).Msg("scan error for list users")
			continue
		}
		result := user.ConvertToPB()
		list = append(list, result)
	}
	err = rows.Err()
	if err != nil {
		log.Error().Err(err).Msg("error in scan multiple values of users")
	}
	return list, nil
}

// GetSession - функция, которая возвращает сессию в структуре protobuf c указанными Login и SessionId
func (r *IdmRepo) GetSession(ctx context.Context, cfg *configs.AppConfig, login string, id string) (*pb.Session, error) {
	sqlBytes, err := SQLSelect.ReadFile("migrations/select/SQLSelectSession.sql")
	if err != nil {
		return nil, err
	}
	sqlQuery := string(sqlBytes)
	row := r.Pool.QueryRow(ctx, sqlQuery, login, id)
	var session PGSession
	err = row.Scan(&session.Id, &session.Login, &session.Config, &session.Created, &session.Duration)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, service.ErrEmptySQLResult
		}
		log.Error().Err(err).Msg("failed get session")
		return nil, err
	}
	result := session.ConvertToPB()
	return result, nil
}

// GetAllSessions - функция, которая возвращает все сессии в структуре protobuf
func (r *IdmRepo) GetAllSessions(ctx context.Context, cfg *configs.AppConfig) ([]*pb.Session, error) {
	var list []*pb.Session
	sqlBytes, err := SQLSelect.ReadFile("migrations/select/SQLSelectAllSessions.sql")
	if err != nil {
		return nil, err
	}
	sqlQuery := string(sqlBytes)
	rows, err := r.Pool.Query(ctx, sqlQuery)
	if err != nil {
		log.Error().Err(err).Msg("execute query failed")
		return list, err
	}
	defer rows.Close()
	for rows.Next() {
		var session PGSession
		err = rows.Scan(&session.Id, &session.Login, &session.Config, &session.Created, &session.Duration)
		if err != nil {
			log.Error().Err(err).Msg("scan error for list sessions")
			continue
		}
		result := session.ConvertToPB()
		list = append(list, result)
	}
	err = rows.Err()
	if err != nil {
		log.Error().Err(err).Msg("error in scan multiple values of sessions")
	}
	return list, nil
}
