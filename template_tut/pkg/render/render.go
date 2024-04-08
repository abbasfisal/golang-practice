package render

import (
	"html/template"
	"log"
	"net/http"
)

func ParseTemplate(w http.ResponseWriter, tmplPath string) {
	tmpl, err := template.ParseFiles(tmplPath, "./templates/base.layout.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}
