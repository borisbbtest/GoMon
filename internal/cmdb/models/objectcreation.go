package models

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/cmdb/service"
	pb "github.com/borisbbtest/GoMon/internal/models/cmdb"
)

// CreateObject - функция создания КЕ в хранилище
func (w *ConfigWrapper) CreateObject(ctx context.Context, ci *pb.Ci) error {
	err := w.Repo.CreateObject(ctx, w.Cfg, ci)
	if err != nil {
		log.Error().Err(err).Msg("failed create ci in db")
		return err
	}
	return nil
}

// CreateBatchObjects - функция создания КЕ в хранилище пакетно (batch)
func (w *ConfigWrapper) CreateBatchObjects(ctx context.Context, cis []*pb.Ci) error {
	errflag := false
	for _, ci := range cis {
		err := w.Repo.CreateObject(ctx, w.Cfg, ci)
		if err != nil {
			log.Error().Err(err).Msg("failed create ci in db")
			errflag = true

		}
	}
	if errflag {
		return service.ErrInsertObjects
	}
	return nil
}
