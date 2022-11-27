package models

import (
	"context"
	"time"

	pb "github.com/borisbbtest/GoMon/internal/models/idm"
)

// CheckAuthorized - метод, возвращают true в случае, если сессия существует в idm и время ее окончания после текущего времени.
// Во всех остальных случаях возвращается false
func (cw *ConfigWrapper) CheckAuthorized(ctx context.Context, login string, id string) bool {
	var req pb.GetSessionRequest
	req.Id = id
	req.Login = login
	resp, err := cw.Conns.Idm.GetSession(ctx, &req)
	if err != nil {
		return false
	}
	ss := resp.Ss
	return time.Now().Before(ss.Duration.AsTime())
}
