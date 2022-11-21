package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/borisbbtest/GoMon/internal/fanout/models"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
)

type WrapperMiddleware struct {
	App *models.ConfigWrapper // структура, хранящая пул подключений и конфиг приложения
}

// CheckAuthorized - middleware для проверки авторизованно ли подключения.
// Получает session_token и login из cookie. Проверяет существует ли сессия и она не протухла.
// Добавляет в контекст поле login c типом models.FanInContextKey для идентификации пользователя в других пакетах.
func (hook *WrapperMiddleware) MiddleSetSessionCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cookiess, err := r.Cookie("session_token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				http.Error(rw, "no cookies session", http.StatusUnauthorized)
				return
			}
			utils.Log.Error().Err(err).Msg("internal error")
			http.Error(rw, "internal errors", http.StatusInternalServerError)
			return
		}
		cookieu, err := r.Cookie("login")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				http.Error(rw, "no cookies login", http.StatusUnauthorized)
				return
			}
			utils.Log.Error().Err(err).Msg("internal error")
			http.Error(rw, "internal errors", http.StatusInternalServerError)
			return
		}
		sessionToken := cookiess.Value
		login := cookieu.Value
		ok := hook.App.CheckAuthorized(r.Context(), login, sessionToken)
		if !ok {
			http.Error(rw, "no session with this session token or session expired", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), models.FanInContextKey("login"), login)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
