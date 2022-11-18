package models

import (
	"context"
	"sync"
	"time"
)

func (w *ConfigWrapper) ClearExpiredWorker(ctx context.Context, wg *sync.WaitGroup) {
	ticker := time.NewTicker(2 * w.Cfg.SessionTimeExpired)
	for {
		select {
		case <-ticker.C:
			err := w.Repo.ClearExpiredSessions(ctx, w.Cfg)
			if err != nil {
				log.Error().Err(err).Msg("failed clear expired sessions")
			}
		case <-ctx.Done():
			ticker.Stop()
			log.Info().Msg("stopped ClearExpiredWorker")
			wg.Done()
			return
		}
	}
}
