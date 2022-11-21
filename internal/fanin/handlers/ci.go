package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/borisbbtest/GoMon/internal/fanin/models"
)

// PushCiHandler - хендлер, описывающий получение КЕ и сохранение его в cmdb
//
// POST [/api/ci/push]
// Входные данные: models.Ci
func (h *HTTP) PushCiHandler(rw http.ResponseWriter, r *http.Request) {
	var ci models.Ci
	if err := json.NewDecoder(r.Body).Decode(&ci); err != nil {
		log.Error().Err(err).Msg("failed decode ci")
		http.Error(rw, "can't decode ci", http.StatusBadRequest)
		return
	}
	err := h.App.PushCi(r.Context(), &ci)
	if err != nil {
		log.Error().Err(err).Msg("failed store ci")
		http.Error(rw, "can't store ci", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "ci saved successfully")
}

// PushBatchCiHandler - хендлер, описывающий получение КЕ пакетно и сохранение его в cmdb
//
// POST [/api/ci/push/batch]
// Входные данные: []models.Ci
func (h *HTTP) PushBatchCiHandler(rw http.ResponseWriter, r *http.Request) {
	var cis []models.Ci
	if err := json.NewDecoder(r.Body).Decode(&cis); err != nil {
		log.Error().Err(err).Msg("failed decode cis")
		http.Error(rw, "can't decode cis", http.StatusBadRequest)
		return
	}
	err := h.App.PushBatchCis(r.Context(), cis)
	if err != nil {
		log.Error().Err(err).Msg("failed store cis")
		http.Error(rw, "can't store cis", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "cis saved successfully")
}

// PushBatchCiHandler - хендлер, описывающий удаление КЕ из cmdb
//
// POST [/api/ci/delete]
// Входные данные: string - имя удаляемой КЕ
func (h *HTTP) DeleteCiHandler(rw http.ResponseWriter, r *http.Request) {
	var ci string
	if err := json.NewDecoder(r.Body).Decode(&ci); err != nil {
		log.Error().Err(err).Msg("failed decode ci")
		http.Error(rw, "can't decode ci", http.StatusBadRequest)
		return
	}
	err := h.App.DeleteCi(r.Context(), ci)
	if err != nil {
		log.Error().Err(err).Msg("failed delete ci")
		http.Error(rw, "can't delete ci", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "ci deleted successfully")
}

// DeleteBatchCiHandler - хендлер, описывающий удаление КЕ пакетно из cmdb
//
// POST [/api/ci/delete/batch]
// Входные данные: []string - массив имен удаляемых КЕ
func (h *HTTP) DeleteBatchCiHandler(rw http.ResponseWriter, r *http.Request) {
	var cis []string
	if err := json.NewDecoder(r.Body).Decode(&cis); err != nil {
		log.Error().Err(err).Msg("failed decode cis")
		http.Error(rw, "can't decode cis", http.StatusBadRequest)
		return
	}
	err := h.App.DeleteBatchCis(r.Context(), cis)
	if err != nil {
		log.Error().Err(err).Msg("failed delete cis")
		http.Error(rw, "can't delete cis", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "cis deleted successfully")
}
