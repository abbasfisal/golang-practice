package handlers

import (
	"golang-practice/template_tut/pkg/config"
	"golang-practice/template_tut/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (re *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.ParseTemplate(w, "./templates/home.page.html")
}

func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.ParseTemplate(w, "./templates/about.page.html")
}
