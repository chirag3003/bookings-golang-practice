package render

import (
	"bytes"
	"git.chirag.codes/chirag/bookings-golang/pkg/config"
	"git.chirag.codes/chirag/bookings-golang/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func addDefaultTemplateData(td *models.TemplateData) *models.TemplateData {
	return td
}
func RenderTemplate(w http.ResponseWriter, tName string, td *models.TemplateData) {
	name := tName + ".pages.tmpl"
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, exists := tc[name]
	if !exists {
		log.Fatal("Template does not exist")
	}
	buf := new(bytes.Buffer)
	err := t.Execute(buf, addDefaultTemplateData(td))
	if err != nil {
		log.Println(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.pages.tmpl")
	if err != nil {
		log.Println(err)
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println(err)
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Println("53", err)
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
				log.Println(err)
			}
		}
		myCache[name] = ts
	}
	return myCache, nil

}
