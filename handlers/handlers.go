package handlers

import (
	"net/http"

	"github.com/nazmulcuet11/bednbreakfast/configs"
	"github.com/nazmulcuet11/bednbreakfast/models"
	"github.com/nazmulcuet11/bednbreakfast/renderers"
)

type Handler struct {
	app *configs.AppConfig
}

func NewHandler(app *configs.AppConfig) *Handler {
	return &Handler{
		app: app,
	}
}

func (h *Handler) Generals(w http.ResponseWriter, r *http.Request) {
	renderers.NewRenderer(h.app).RenderTemplate(w, "generals.page.gohtml", &models.TemplateData{})
}

func (h *Handler) Majors(w http.ResponseWriter, r *http.Request) {
	renderers.NewRenderer(h.app).RenderTemplate(w, "majors.page.gohtml", &models.TemplateData{})
}

func (h *Handler) SearchAvailablity(w http.ResponseWriter, r *http.Request) {
	renderers.NewRenderer(h.app).RenderTemplate(w, "search-availability.page.gohtml", &models.TemplateData{})
}

func (h *Handler) Contact(w http.ResponseWriter, r *http.Request) {
	renderers.NewRenderer(h.app).RenderTemplate(w, "contact.page.gohtml", &models.TemplateData{})
}

func (h *Handler) MakeReservation(w http.ResponseWriter, r *http.Request) {
	renderers.NewRenderer(h.app).RenderTemplate(w, "make-reservation.page.gohtml", &models.TemplateData{})
}
