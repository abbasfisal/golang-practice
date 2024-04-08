package main

import (
	"golang-practice/template_tut/pkg/config"
	"golang-practice/template_tut/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Println("star serve")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("server: failed ", err.Error())
		return
	}
}
