package api

import (
	"net/http"
	"reflect"
	"runtime"

	"github.com/kwanifi/numiscan-api/config"
	"github.com/kwanifi/numiscan-api/dao/filters"
	"github.com/kwanifi/numiscan-api/log"
	"github.com/kwanifi/numiscan-api/smodels"
)

func (api *API) Index(w http.ResponseWriter, r *http.Request) {
	jsonData(w, map[string]string{
		"service": config.ServiceName,
	})
}

func (api *API) Health(w http.ResponseWriter, r *http.Request) {
	jsonData(w, map[string]bool{
		"status": true,
	})
}

func (api *API) aggHandler(w http.ResponseWriter, r *http.Request, action func(filters.Agg) ([]smodels.AggItem, error)) {
	method := runtime.FuncForPC(reflect.ValueOf(action).Pointer()).Name()
	var filter filters.Agg
	err := api.queryDecoder.Decode(&filter, r.URL.Query())
	if err != nil {
		log.Debug("API %s: Decode: %s", method, err.Error())
		jsonBadRequest(w, "")
		return
	}
	err = filter.Validate()
	if err != nil {
		log.Debug("API %s: Validate: %s", method, err.Error())
		jsonBadRequest(w, err.Error())
		return
	}
	resp, err := action(filter)
	if err != nil {
		log.Error("API %s: %s", method, err.Error())
		jsonError(w)
		return
	}
	jsonData(w, resp)
}
