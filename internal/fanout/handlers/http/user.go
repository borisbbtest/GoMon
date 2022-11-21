package handlers_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	integrationidm "github.com/borisbbtest/GoMon/internal/fanout/clients/grpc/idm"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
)

// RegisterHandler - хендлер, описывающий регистрацию нового пользователя, сохранение его в idm и создание сессии для него.
// Добавляет 2 cookie: session_token и login.
//
// POST [/api/register]
// Входные данные: models.User
func (hook *WrapperHandler) RegisterHandler(rw http.ResponseWriter, r *http.Request) {
	var user integrationidm.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Log.Error().Err(err).Msg("failed decode user")
		http.Error(rw, "can't decode user", http.StatusBadRequest)
		return
	}
	session, err := hook.ServicePool.Idm.RegisterUser(r.Context(), &user)
	if err != nil {
		utils.Log.Error().Err(err).Msg("failed register user")
		http.Error(rw, "failed register user", http.StatusInternalServerError)
		return
	}
	cookieSession := &http.Cookie{
		Name:    "session_token",
		Value:   session.Id,
		Expires: session.Duration,
	}
	cookieUser := &http.Cookie{
		Name:    "login",
		Value:   user.Login,
		Expires: session.Duration,
	}
	http.SetCookie(rw, cookieSession)
	http.SetCookie(rw, cookieUser)
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "user registered successfully")
}

// AuthorizeHandler - хендлер, авторизации пользователя и создания для него сессии.
//
// POST [/api/authorize]
// Входные данные:
//
//	{
//		"login": "username",
//		"password": "password"
//	}
func (hook *WrapperHandler) AuthorizeHandler(rw http.ResponseWriter, r *http.Request) {
	var user integrationidm.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Log.Error().Err(err).Msg("failed decode user")
		http.Error(rw, "can't decode user", http.StatusBadRequest)
		return
	}
	session, err := hook.ServicePool.Idm.AuthorizeUser(r.Context(), &user)
	if err != nil {
		utils.Log.Error().Err(err).Msg("failed authorize user")
		http.Error(rw, "failed authorize user", http.StatusInternalServerError)
		return
	}
	cookieSession := &http.Cookie{
		Name:    "session_token",
		Value:   session.Id,
		Expires: session.Duration,
	}
	cookieUser := &http.Cookie{
		Name:    "login",
		Value:   user.Login,
		Expires: session.Duration,
	}
	http.SetCookie(rw, cookieSession)
	http.SetCookie(rw, cookieUser)
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "user authorize successfully")
}
