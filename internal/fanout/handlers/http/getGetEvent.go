package handlers_http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/borisbbtest/GoMon/internal/fanout/models"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
	"github.com/go-chi/chi/v5"
)

// GetGetEvent - получаем одно событие оп uuid
func (hook *WrapperHandler) GetGetEvent(w http.ResponseWriter, r *http.Request) {
	utils.Log.Info().Msg("GetGetEvent")
	w.Header().Add("Content-Type", "application/json")

	uuid := chi.URLParam(r, "uuid")
	utils.Log.Info().Msg(uuid)
	defer r.Body.Close()

	resgrpc, err := hook.ServicePool.Mgrevent.GetEvent(r.Context(), uuid)
	//utils.Log.Debug().Msgf(" ---   %s", test)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	resp := models.ResponseGetEvent{Root: resgrpc}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
		utils.Log.Error().Err(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Post handler")
}
