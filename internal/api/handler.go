package api

import (
	"encoding/json"
	"net/http"
	"sweetweather.test/pkg/response"
)

type CalucalateBody struct {
	Data string
}

func (s *Server) HandleCalculate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body CalucalateBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.WriteErrResponse(w, "bad request")
			return
		}

		res, err := s.uc.Calculate(body.Data)
		if err != nil {
			response.WriteErrResponse(w, err.Error())
			return
		}

		response.WriteOkResponse(w, struct {
			Data int `json:"data"`
		}{
			Data: res,
		})
	}
}
