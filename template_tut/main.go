package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	log.Println("star serve")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("server: failed ", err.Error())
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	parseTemplate(w, "./templates/home.page.tmpl")
}

func about(w http.ResponseWriter, r *http.Request) {
	parseTemplate(w, "./templates/about.page.html")
}

func parseTemplate(w http.ResponseWriter, tmplPath string) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}
