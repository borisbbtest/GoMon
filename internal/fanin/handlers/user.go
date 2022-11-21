package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/borisbbtest/GoMon/internal/fanin/models"
)

// RegisterHandler - хендлер, описывающий регистрацию нового пользователя, сохранение его в idm и создание сессии для него.
// Добавляет 2 cookie: session_token и login.
//
// POST [/api/register]
// Входные данные: models.User
func (h *HTTP) RegisterHandler(rw http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Error().Err(err).Msg("failed decode user")
		http.Error(rw, "can't decode user", http.StatusBadRequest)
		return
	}
	session, err := h.App.RegisterUser(r.Context(), &user)
	if err != nil {
		log.Error().Err(err).Msg("failed register user")
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
func (h *HTTP) AuthorizeHandler(rw http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Error().Err(err).Msg("failed decode user")
		http.Error(rw, "can't decode user", http.StatusBadRequest)
		return
	}
	session, err := h.App.AuthorizeUser(r.Context(), &user)
	if err != nil {
		log.Error().Err(err).Msg("failed authorize user")
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
