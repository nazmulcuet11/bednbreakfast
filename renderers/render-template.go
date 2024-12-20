package renderers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/nazmulcuet11/bednbreakfast/models"
)

// RenderTemplate renders a template
func (r *Renderer) RenderTemplate(w http.ResponseWriter, page string, td *models.TemplateData) {
	tmpl, err := r.getTemplate(page)
	if err != nil {
		r.app.ErrorLogger.Fatal(err)
	}

	td = r.addDefaultData(td)
	err = tmpl.Execute(w, td)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func (r *Renderer) getTemplate(page string) (*template.Template, error) {
	if r.app.UseTemplateCache {
		return r.getTemplateThroughCache(page)
	} else {
		return r.getTemplateFromDisk(page)
	}
}

func (r *Renderer) getTemplateThroughCache(page string) (*template.Template, error) {
	if cachedTempl, ok := r.app.TemplateCache[page]; ok {
		return cachedTempl, nil
	} else {
		return r.getTemplateFromDiskAndCache(page)
	}
}

func (r *Renderer) getTemplateFromDiskAndCache(page string) (*template.Template, error) {
	tmpl, err := r.getTemplateFromDisk(page)
	if err != nil {
		return nil, err
	}
	r.app.TemplateCache[page] = tmpl
	return tmpl, nil
}

func (r *Renderer) getTemplateFromDisk(page string) (*template.Template, error) {
	t, err := template.ParseFiles("./templates/" + page)
	t, err = t.ParseGlob("./templates/*.layout.gohtml")
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *Renderer) addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
