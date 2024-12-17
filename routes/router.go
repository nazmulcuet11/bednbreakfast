package routes

import "github.com/nazmulcuet11/bednbreakfast/configs"

type Router struct {
	app *configs.AppConfig
}

func NewRouter(app *configs.AppConfig) *Router {
	return &Router{
		app: app,
	}
}
