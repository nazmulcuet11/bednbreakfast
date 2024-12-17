package middlewares

import "github.com/nazmulcuet11/bednbreakfast/configs"

type Middleware struct {
	app *configs.AppConfig
}

func NewMiddleware(app *configs.AppConfig) *Middleware {
	return &Middleware{
		app: app,
	}
}
