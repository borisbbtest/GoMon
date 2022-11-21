package handlers_http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/borisbbtest/GoMon/internal/fanout/models"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
)

// PostGetCis получаем список КЕ
func (hook *WrapperHandler) PostGetCis(w http.ResponseWriter, r *http.Request) {

	utils.Log.Info().Msg("PostGetCis")
	w.Header().Add("Content-Type", "application/json")

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	defer r.Body.Close()
	req := models.RequestGetCis{}
	req.ParseRequest(bytes)

	resgrpc, err := hook.ServicePool.Cmdb.GetBachListItem(r.Context(), req.Item)
	//utils.Log.Debug().Msgf(" ---   %s", test)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	resp := models.ResponseGetCis{Root: resgrpc}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
		utils.Log.Error().Err(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Post handler")
}
