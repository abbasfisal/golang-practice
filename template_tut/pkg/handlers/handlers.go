package handlers

import (
	"golang-practice/template_tut/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.ParseTemplate(w, "./templates/home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.ParseTemplate(w, "./templates/about.page.html")
}
