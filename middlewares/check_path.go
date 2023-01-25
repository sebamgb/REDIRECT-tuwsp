package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	mem "tuwsp/memory"
	"tuwsp/server"
)

type checkResponse struct {
	UserInfo string   `json:"user_info"`
	UrlInfo  []string `json:"url_info"`
}

// CheckMiddleware check the path
func CheckPathMiddleware(s server.Server) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// declare wait group to count and wait for concurrence
			var wg sync.WaitGroup
			// making cache
			get := mem.NewGetSomethingAdapter(r.Context(), false)
			cache := mem.NewCache(get.GetSomething)
			// cutting uri
			path := strings.
				Join(strings.Split(r.RequestURI, "/"), "")
			// parameters to func main
			m := map[string]string{
				"user":        path,
				"protocol_id": "s",
			}
			// translate to json
			dataByte, err := json.Marshal(m)
			if err != nil {
				return
			}
			// declare info
			var info []byte
			// increment count
			wg.Add(1)
			// goroutine for cache concurrent
			go func(w http.ResponseWriter) {
				// decrement count
				defer wg.Done()
				// main func with json
				info, err = cache.Get(string(dataByte))
				if err != nil {
					http.Redirect(w, r, "/", http.StatusFound)
					return
				}
			}(w)
			// wait for concurrence
			wg.Wait()
			// translate from json byte
			var response checkResponse
			err = json.Unmarshal(info, &response)
			if err != nil {
				return
			}
			// auth to header
			var auth string
			auth = response.UserInfo
			// getting and setting header for  user
			authUser := w.Header().Get("Authorization")
			if authUser == "" {
				authUser = auth
			}
			w.Header().Set("Authorization", authUser)
			// getting and setting header for url
			url := w.Header().Get("X-Url-Id")
			if url == "" {
				bytes, err := json.Marshal(response.UrlInfo)
				if err != nil {
					log.Fatal(err)
				}
				url = string(bytes)
			}
			w.Header().Set("X-Url-Id", url)
			// continue
			h.ServeHTTP(w, r)
		})
	}
}
