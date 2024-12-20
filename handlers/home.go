package handlers

import (
	"net/http"

	"github.com/nazmulcuet11/bednbreakfast/models"
	"github.com/nazmulcuet11/bednbreakfast/renderers"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	h.app.SessionManager.Put(r.Context(), "remote_ip", remoteIP)
	renderers.NewRenderer(h.app).RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}
