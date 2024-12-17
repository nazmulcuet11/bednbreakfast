package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/nazmulcuet11/bednbreakfast/configs"
	"github.com/nazmulcuet11/bednbreakfast/routes"
)

func main() {
	app := configs.AppConfig{
		PortNumber:       ":8080",
		IsProduction:     false, // read from env
		UseTemplateCache: false, // read from env
		TemplateCache:    make(map[string]*template.Template),
		InfoLogger:       log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime),
		ErrorLogger:      log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime),
	}

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.IsProduction

	app.SessionManager = sessionManager

	srv := &http.Server{
		Addr:    app.PortNumber,
		Handler: routes.NewRouter(&app).Routes(),
	}

	app.InfoLogger.Printf("starting server at: %s\n", app.PortNumber)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
