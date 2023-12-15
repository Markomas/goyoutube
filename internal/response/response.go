package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	data  interface{}
	error error
}

func ErrorResponse(w http.ResponseWriter, err error, status int) {
	resp := Response{error: err}
	respJson, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(respJson), status)
}
