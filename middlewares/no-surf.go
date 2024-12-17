package middlewares

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func (m *Middleware) NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   m.app.IsProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
