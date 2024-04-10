package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	http.HandleFunc("/tmp", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("about.html", "base.html")
		data := map[string]interface{}{
			"title": "about us title",
		}
		if err != nil {
			log.Fatal(err)
			return
		}
		t.Execute(w, data)
		return
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		films := map[string][]Film{
			"Films": {
				{
					Title:    "The God Father",
					Director: "jimmy",
				},
				{
					Title:    "Blade Runner",
					Director: "Momnii",
				},
			},
		}

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, films)
	})
	log.Println("\n start server")
	http.ListenAndServe(":8080", nil)
}
