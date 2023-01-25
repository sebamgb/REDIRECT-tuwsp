package middlewares

import (
	"net/http"
	"sync"

	jsontuwsp "tuwsp/json_tuwsp"
	"tuwsp/server"
)

func InsertJsonMiddleware(s server.Server) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var once = &sync.Once{}
			go func() {
				once.Do(func() {
					err := jsontuwsp.InsertProtocolJson()
					if err != nil {
						return
					}
					err = jsontuwsp.InsertUrlJson()
					if err != nil {
						return
					}
					err = jsontuwsp.InsertUserJson()
					if err != nil {
						return
					}
					err = jsontuwsp.InsertInfoUserJson()
					if err != nil {
						return
					}
					err = jsontuwsp.InsertQueryKeyJson()
					if err != nil {
						return
					}
					err = jsontuwsp.InsertQueryValueJson()
					if err != nil {
						return
					}
					err = jsontuwsp.InsertEndpointJson()
					if err != nil {
						return
					}
				})
			}()
			h.ServeHTTP(w, r)
		})
	}
}
