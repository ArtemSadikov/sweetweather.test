package response

import (
	"encoding/json"
	"net/http"
)

func WriteErrResponse(w http.ResponseWriter, msg string) {
	res, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: msg,
	})
	w.Header().Add("Content-type", "application/json")
	_, _ = w.Write(res)
}
