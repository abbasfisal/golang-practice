package main

import (
	"golang-practice/template_tut/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("star serve")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("server: failed ", err.Error())
		return
	}
}
