package handlers

import "github.com/nazmulcuet11/bednbreakfast/configs"

type Handler struct {
	app *configs.AppConfig
}

func NewHandler(app *configs.AppConfig) *Handler {
	return &Handler{
		app: app,
	}
}
