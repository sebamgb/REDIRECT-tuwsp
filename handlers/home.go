package handlers

import (
	"net/http"
	"tuwsp/server"
)

func RootHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.Config().Debug {
			http.ServeFile(w, r, "development.html")
		}
	}
}
