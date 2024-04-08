package homecontroller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("src/views/home/index.html")
	tmpl.Execute(w, nil)
}
