package postgres

import (
	"context"
	"embed"
	"errors"

	"github.com/borisbbtest/GoMon/internal/idm/configs"
	pb "github.com/borisbbtest/GoMon/internal/idm/proto/idm"
	"github.com/borisbbtest/GoMon/internal/idm/service"
	"github.com/jackc/pgx/v4"
)

//go:embed migrations/select/*.sql
var SQLSelect embed.FS

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

func (r *IdmRepo) GetSession(ctx context.Context, cfg *configs.AppConfig, login string, id string) (*pb.Session, error) {
	sqlBytes, err := SQLSelect.ReadFile("migrations/select/SQLSelectSession.sql")
	if err != nil {
		return nil, err
	}
	sqlQuery := string(sqlBytes)
	row := r.Pool.QueryRow(ctx, sqlQuery, login)
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
