package handler

import (
	"encoding/json"
	"net/http"

	"github.com/decoch/dojo4/kadai4/decoch/application"
)

type OmikujiHandler struct {
	Service application.OmikujiService
}

func (o *OmikujiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")
	encoder.Encode(o.Service.Get())
}
