package configs

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	PortNumber       string
	IsProduction     bool
	UseTemplateCache bool
	TemplateCache    map[string]*template.Template
	SessionManager   *scs.SessionManager
	InfoLogger       *log.Logger
	ErrorLogger      *log.Logger
}
