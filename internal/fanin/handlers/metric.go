package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/borisbbtest/GoMon/internal/fanin/models"
)

// PushMetricHandler - хендлер, описывающий получение метрики и сохранение ее в metrics
//
// POST [/api/metric/push]
// Входные данные: models.Metric
func (h *HTTP) PushMetricHandler(rw http.ResponseWriter, r *http.Request) {
	var metric models.Metric
	if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		log.Error().Err(err).Msg("failed decode metric")
		http.Error(rw, "can't decode metric", http.StatusBadRequest)
		return
	}
	err := h.App.PushMetrics(r.Context(), &metric)
	if err != nil {
		log.Error().Err(err).Msg("failed store metric")
		http.Error(rw, "can't store metric", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "metric saved successfully")
}

// PushBatchMetricHandler - хендлер, описывающий получение метрик пакетно и сохранение их в metrics
//
// POST [/api/metric/push/batch]
// Входные данные: []models.Metric
func (h *HTTP) PushBatchMetricHandler(rw http.ResponseWriter, r *http.Request) {
	var metrics []models.Metric
	if err := json.NewDecoder(r.Body).Decode(&metrics); err != nil {
		log.Error().Err(err).Msg("failed decode metrics")
		http.Error(rw, "can't decode metrics", http.StatusBadRequest)
		return
	}
	err := h.App.PushBatchMetrics(r.Context(), metrics)
	if err != nil {
		log.Error().Err(err).Msg("failed store metrics")
		http.Error(rw, "can't store metrics", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "cis saved successfully")
}
