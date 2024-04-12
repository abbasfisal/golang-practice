package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/getform", getFormHandler)
	http.HandleFunc("/processget", processGetHandler)
	http.ListenAndServe(":8080", nil)
}

func processGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processing form req")
	var f Form
	f.UserName = r.FormValue("username")
	f.Data = r.FormValue("data")

	t, _ := template.ParseFiles("templates/thanks.html")
	t.Execute(w, f)
}

type Form struct {
	UserName string
	Data     string
}

func getFormHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/getform.html")
	t.Execute(w, nil)
}
