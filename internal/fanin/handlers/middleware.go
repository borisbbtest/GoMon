package handlers

import (
	"compress/gzip"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/borisbbtest/GoMon/internal/fanin/models"
)

// gzipWriter - новый writer для использования с gzip
type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

// Write - реализация интерфейса Writer
func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// GzipHandle - middleware для подмены writer на другой с использованием gzip. Устанавливает header "Content-Encoding" на "gzip"
func (h *HTTP) GzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(rw, r)
			return
		}
		gz, err := gzip.NewWriterLevel(rw, gzip.BestSpeed)
		if err != nil {
			io.WriteString(rw, err.Error())
			return
		}
		defer gz.Close()
		rw.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzipWriter{ResponseWriter: rw, Writer: gz}, r)
	})
}

// CheckAuthorized - middleware для проверки авторизованно ли подключения.
// Получает session_token и login из cookie. Проверяет существует ли сессия и она не протухла.
// Добавляет в контекст поле login c типом models.FanInContextKey для идентификации пользователя в других пакетах.
func (h *HTTP) CheckAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cookiess, err := r.Cookie("session_token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				http.Error(rw, "no cookies session", http.StatusUnauthorized)
				return
			}
			log.Error().Err(err).Msg("internal error")
			http.Error(rw, "internal errors", http.StatusInternalServerError)
			return
		}
		cookieu, err := r.Cookie("login")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				http.Error(rw, "no cookies login", http.StatusUnauthorized)
				return
			}
			log.Error().Err(err).Msg("internal error")
			http.Error(rw, "internal errors", http.StatusInternalServerError)
			return
		}
		sessionToken := cookiess.Value
		login := cookieu.Value
		ok := h.App.CheckAuthorized(r.Context(), login, sessionToken)
		if !ok {
			http.Error(rw, "no session with this session token or session expired", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), models.FanInContextKey("login"), login)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
