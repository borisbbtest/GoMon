package models

import (
	"context"
)

// DeleteUser - функция удаления пользователя из хранилища
func (w *ConfigWrapper) DeleteUser(ctx context.Context, login string) error {
	err := w.Repo.DeleteUser(ctx, w.Cfg, login)
	if err != nil {
		log.Error().Err(err).Msg("failed delete user")
		return err
	}
	return nil
}
