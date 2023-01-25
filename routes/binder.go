package routes

import (
	"tuwsp/handlers"
	"tuwsp/middlewares"
	"tuwsp/server"

	"github.com/gorilla/mux"
)

// BindRoute implement Router mux of gorilla for call
func BindRoute(s server.Server, r *mux.Router) {
	public := r.NewRoute().Subrouter()

	public.Use(middlewares.InsertJsonMiddleware(s))
	public.Use(middlewares.CheckPathMiddleware(s))
	public.Path("/{[a-z]+}").
		Handler(handlers.RedirectHandler(s))
	r.Handle("/", handlers.RootHandler(s))
}
