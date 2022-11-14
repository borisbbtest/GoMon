package models

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/cmdb/service"
	pb "github.com/borisbbtest/GoMon/internal/models/cmdb"
)

// GetObject - функция получения КЕ из хранилища
func (w *ConfigWrapper) GetObject(ctx context.Context, ci string) (*pb.Ci, error) {
	user, err := w.Repo.GetObject(ctx, w.Cfg, ci)
	if err != nil {
		log.Error().Err(err).Msg("failed get ci")
		return nil, err
	}
	return user, nil
}

// GetBatchObjects - функция получения КЕ из хранилища.
// В случае ошибки получения какой-либо КЕ из списка, она пропускается и возвращается ошибка.
// В случае возвращения ошибки так же возвращается массив успешно полученных КЕ.
func (w *ConfigWrapper) GetBatchObjects(ctx context.Context, cis []string) ([]*pb.Ci, error) {
	errflag := false
	var result []*pb.Ci
	for _, ci := range cis {
		user, err := w.Repo.GetObject(ctx, w.Cfg, ci)
		if err != nil {
			log.Error().Err(err).Msg("failed get ci")
			errflag = true
		}
		result = append(result, user)
	}
	if errflag {
		return result, service.ErrSelectObjects
	}
	return result, nil
}
