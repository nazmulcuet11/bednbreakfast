package middlewares

import (
	"net/http"
)

func (m *Middleware) SessionLoad(next http.Handler) http.Handler {
	return m.app.SessionManager.LoadAndSave(next)
}
