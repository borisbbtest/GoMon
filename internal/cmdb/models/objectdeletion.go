package models

import (
	"context"

	"github.com/borisbbtest/GoMon/internal/cmdb/service"
)

// DeleteObject - функция удаления КЕ из хранилища
func (w *ConfigWrapper) DeleteObject(ctx context.Context, name string) error {
	err := w.Repo.DeleteObject(ctx, w.Cfg, name)
	if err != nil {
		log.Error().Err(err).Msg("failed delete Ci")
		return err
	}
	return nil
}

// DeleteBatchObject - функция удаления КЕ из хранилища пакетно
func (w *ConfigWrapper) DeleteBatchObject(ctx context.Context, cis []string) error {
	errflag := false
	for _, ci := range cis {
		err := w.Repo.DeleteObject(ctx, w.Cfg, ci)
		if err != nil {
			log.Error().Err(err).Msg("failed delete Ci")
			errflag = true
		}
	}
	if errflag {
		return service.ErrDeleteObjects
	}
	return nil
}
