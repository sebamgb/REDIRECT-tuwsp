package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	mem "tuwsp/memory"
	"tuwsp/pattern"
	"tuwsp/server"
)

type redirectResponse struct {
	UrlPath string `json:"url_path"`
}

// RedirectHandler if BuildUrl returns an url then redirect
func RedirectHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// declare wait group for count and wait for concurrence
		var wg sync.WaitGroup
		// reception from checkpath middleware
		authUser := w.Header().Get("Authorization")
		urlAuth := w.Header().Get("X-Url-Id")
		// preparing cache
		get := mem.NewGetSomethingAdapter(r.Context(), true)
		cache := mem.NewCache(get.GetSomething)
		// urlAuth unmarshal
		var urlHeader []string
		json.Unmarshal([]byte(urlAuth), &urlHeader)
		// make userId
		sb := strings.Builder{}
		sb.WriteString(authUser)
		sb.WriteString("_")
		sb.WriteString(urlHeader[1])
		userId := sb.String()
		// preparing parameters
		m := map[string]string{
			"user":     authUser,
			"user_id":  userId,
			"domain":   urlHeader[0],
			"url_id":   urlHeader[1],
			"protocol": urlHeader[2],
		}
		// translate parameters to json bytes
		jsonBytes, err := json.Marshal(m)
		if err != nil {
			return
		}
		// declare infom as bytes
		var info []byte
		// increment count
		wg.Add(1)
		// goroutine for cache concurrent
		go func(w http.ResponseWriter) {
			// decrement count
			defer wg.Done()
			// wrap parameters to main func
			info, err = cache.Get(string(jsonBytes))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}(w)
		// wait for concurrence
		wg.Wait()
		// instance response
		var response redirectResponse
		// translate from json byte to response
		err = json.Unmarshal(info, &response)
		if err != nil {
			return
		}
		// observer from uri
		redirect := pattern.NewRedirect(authUser)
		// show path
		request := NewHandler(response.UrlPath)
		// registing path
		redirect.Register(request)
		// compare with user authentic
		redirect.UpdateFigureIn(authUser)
		if redirect.Figure() {
			http.Redirect(w, r, response.UrlPath, http.StatusSeeOther)
		}
	}
}
