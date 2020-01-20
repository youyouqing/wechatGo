package util

import (
	"encoding/json"
	"net/http"
)


func ResultJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(data)
	w.Write(json)
}
