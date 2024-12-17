package handlers

import (
	"net/http"

	"github.com/nazmulcuet11/bednbreakfast/models"
	"github.com/nazmulcuet11/bednbreakfast/renderers"
)

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]interface{})
	stringMap["test"] = "Hello, again"

	remoteIP := h.app.SessionManager.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	renderers.NewRenderer(h.app).RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		Data: stringMap,
	})
}
