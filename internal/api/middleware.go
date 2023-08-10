package api

import (
	"net/http"
	"sweetweather.test/pkg/response"
)

func authMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("User-Access")

		if header != "superuser" {
			response.WriteErrResponse(w, "unauthorized")
			return
		}

		h.ServeHTTP(w, r)
	})
}
