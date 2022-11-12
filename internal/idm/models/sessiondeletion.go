package models

import (
	"context"
)

// DeleteSession - функция удаления сессии из хранилища
func (w *ConfigWrapper) DeleteSession(ctx context.Context, login string, id string) error {
	err := w.Repo.DeleteSession(ctx, w.Cfg, login, id)
	if err != nil {
		log.Error().Err(err).Msg("failed delete session")
		return err
	}
	return nil
}
