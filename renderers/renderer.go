package renderers

import (
	"github.com/nazmulcuet11/bednbreakfast/configs"
)

type Renderer struct {
	app *configs.AppConfig
}

func NewRenderer(app *configs.AppConfig) *Renderer {
	return &Renderer{
		app: app,
	}
}
