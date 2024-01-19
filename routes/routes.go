package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start_Routes() (r *chi.Mux) {
	r = chi.NewRouter()
	def_middleware(r)
	def_routes(r)
	return
}

func def_middleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}

func def_routes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "tmpl/index.html")
	})
	r.Get("/first", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "tmpl/first.html")
	})
	r.Get("/second", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "tmpl/second.html")
	})
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
