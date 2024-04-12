package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("1nd-cookie")
	if err != nil {
		fmt.Println("cookie not found")
		cookie = &http.Cookie{
			Name:     "1nd-cookie",
			Value:    "my 1nd cookie value",
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie.Name)
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}
