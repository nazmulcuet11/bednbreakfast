package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nazmulcuet11/bednbreakfast/handlers"

	"github.com/nazmulcuet11/bednbreakfast/middlewares"
)

func (r *Router) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middlewares.NewMiddleware(r.app).NoSurf)
	mux.Use(middlewares.NewMiddleware(r.app).SessionLoad)

	mux.Get("/", handlers.NewHandler(r.app).Home)
	mux.Get("/about", handlers.NewHandler(r.app).About)
	mux.Get("/generals-quarters", handlers.NewHandler(r.app).Generals)
	mux.Get("/majors-suite", handlers.NewHandler(r.app).Majors)
	mux.Get("/search-availability", handlers.NewHandler(r.app).SearchAvailablity)
	mux.Post("/search-availability", handlers.NewHandler(r.app).PostSearchAvailablity)
	mux.Get("/contact", handlers.NewHandler(r.app).Contact)
	mux.Get("/make-reservation", handlers.NewHandler(r.app).MakeReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
