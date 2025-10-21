package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, data any) {
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
