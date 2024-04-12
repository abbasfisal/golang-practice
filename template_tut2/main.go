package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", Index)
	log.Println("start server : 8080")
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	var tmpl, err = template.New("index.html").Funcs(template.FuncMap{
		"Hi": func() string {
			return "hi user"
		},
		"HiUser": func() string {
			return SayHi()
		},
	}).ParseFiles("index.html")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	tmpl.Execute(w, tmpl)
}

func SayHi() string {
	return "hi and welcome"
}
