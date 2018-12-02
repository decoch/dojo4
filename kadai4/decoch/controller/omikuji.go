package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/decoch/dojo4/kadai4/decoch/model"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")
	encoder.Encode(model.DrawFortuneSlip(time.Now()))
}
