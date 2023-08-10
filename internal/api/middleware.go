package api

import (
	"log"
	"net/http"
	"sweetweather.test/pkg/response"
)

func authMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("User-Access")

		if header != "superuser" {
			msg := "unauthorized"
			log.Print(msg)
			response.WriteErrResponse(w, msg)
			return
		}

		h.ServeHTTP(w, r)
	})
}
