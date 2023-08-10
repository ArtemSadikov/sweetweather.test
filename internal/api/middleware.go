package api

import (
	"encoding/json"
	"net/http"
)

func authMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("User-Access")

		if header != "superuser" {
			res, _ := json.Marshal(struct{ message string }{
				message: "unauthorized",
			})
			w.Header().Add("Content-type", "application/json")
			_, _ = w.Write(res)
			return
		}

		h.ServeHTTP(w, r)
	})
}
