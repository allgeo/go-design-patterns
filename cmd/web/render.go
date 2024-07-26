package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type templateDate struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, t string, td *templateDate) {
	var tmpl *template.Template

	// if we are using template cache, try to ge tit from out map, stored in the recever
	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[t]; ok {
			tmpl = templateFromMap
		}
	}

	// if there is no template in cache,
	if tmpl == nil {
		newTemplate, err := app.buildTemplateFromDisk(t)
		if err != nil {
			log.Println("Error building template: ", err)
		}
		log.Println("building template form disk")
		tmpl = newTemplate
	}

	if td == nil {
		td = &templateDate{}
	}

	if err := tmpl.ExecuteTemplate(w, t, td); err != nil {
		log.Println("Error executing template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (app *application) buildTemplateFromDisk(t string) (*template.Template, error) {
	templateSlice := []string{
		"./templates/base.tmpl",
		"./templates/header.tmpl",
		"./templates/footer.tmpl",
		fmt.Sprintf("./templates/%s", t),
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}
	app.templateMap[t] = tmpl

	return tmpl, nil

}
