package handlers_http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/borisbbtest/GoMon/internal/fanout/models"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
)

// PostGetMetrics - получаем метрики по http
func (hook *WrapperHandler) PostGetMetrics(w http.ResponseWriter, r *http.Request) {

	utils.Log.Info().Msg("PostGetMetrics")
	w.Header().Add("Content-Type", "application/json")

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	defer r.Body.Close()
	req := models.RequestGetMetricDuration{}
	req.ParseRequest(bytes)

	resgrpc, err := hook.ServicePool.Metrics.GetMetricDuration(r.Context(), req.StartTime, req.EndTime)
	//utils.Log.Debug().Msgf(" ---   %s", test)
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	resp := models.ResponseGetMetricDuration{Root: resgrpc}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
		utils.Log.Error().Err(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Post handler")

}
