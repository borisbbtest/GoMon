package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/borisbbtest/GoMon/internal/fanin/models"
)

// PushEventHandler - хендлер, описывающий получение события и сохранение его в events
//
// POST [/api/event/push]
// Входные данные: models.Event
func (h *HTTP) PushEventHandler(rw http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		log.Error().Err(err).Msg("failed decode event")
		http.Error(rw, "can't decode event", http.StatusBadRequest)
		return
	}
	err := h.App.PushEvent(r.Context(), &event)
	if err != nil {
		log.Error().Err(err).Msg("failed store event")
		http.Error(rw, "can't store event", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "event saved successfully")
}

// PushEventHandler - хендлер, описывающий получение событий пакетно и сохранение их в events
//
// POST [/api/event/push/batch]
// Входные данные: []models.Event
func (h *HTTP) PushBatchEventHandler(rw http.ResponseWriter, r *http.Request) {
	var events []models.Event
	if err := json.NewDecoder(r.Body).Decode(&events); err != nil {
		log.Error().Err(err).Msg("failed decode events")
		http.Error(rw, "can't decode events", http.StatusBadRequest)
		return
	}
	err := h.App.PushBatchEvents(r.Context(), events)
	if err != nil {
		log.Error().Err(err).Msg("failed store events")
		http.Error(rw, "can't store events", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "cis saved successfully")
}
