package response

import (
	"encoding/json"
	"net/http"
)

func WriteOkResponse(w http.ResponseWriter, value interface{}) {
	if value == nil {
		WriteErrResponse(w, "internal")
		return
	}
	res, _ := json.Marshal(value)
	w.Header().Add("Content-type", "application/json")
	_, _ = w.Write(res)
	return
}
