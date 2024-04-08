package render

import (
	"html/template"
	"log"
	"net/http"
)

func ParseTemplate(w http.ResponseWriter, tmplPath string) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}
