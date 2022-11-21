package handlers_http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/borisbbtest/GoMon/internal/fanout/models"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
	"github.com/go-chi/chi/v5"
)

// GetGetCi - получаем  одну КЕ по запросу
func (hook *WrapperHandler) GetGetCi(w http.ResponseWriter, r *http.Request) {
	utils.Log.Info().Msg("GetGetCi")
	w.Header().Add("Content-Type", "application/json")

	name := chi.URLParam(r, "name")
	utils.Log.Info().Msg(name)
	defer r.Body.Close()

	resgrpc, err := hook.ServicePool.Cmdb.GetItem(r.Context(), &name)
	//utils.Log.Debug().Msgf(" ---   %s", test)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	resp := models.ResponseGetCi{Root: resgrpc}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
		utils.Log.Error().Err(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Post handler")
}
