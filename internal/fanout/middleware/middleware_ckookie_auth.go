package middleware

import (
	"net/http"

	"github.com/borisbbtest/GoMon/internal/fanout/storage"
)

type WrapperMiddleware struct {
	Session *storage.SessionHTTP
}

func (hook *WrapperMiddleware) MiddleSetSessionCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if true {
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Login failed"))
	})
}
