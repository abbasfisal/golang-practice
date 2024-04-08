package logincontroller

import (
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("src/views/admin/login/index.html")
	tmpl.Execute(w, nil)
}

func Process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	log.Println(username, password, "hi")

	if username == "ali" && password == "123" {
		session, _ := store.Get(r, "mysession")
		session.Values["username"] = username
		sessions.Save(r, w)

		http.Redirect(w, r, "/admin/category", http.StatusSeeOther)
	} else {

		data := map[string]interface{}{
			"error": "invalid username|password",
		}
		tmpl, _ := template.ParseFiles("src/views/admin/login/index.html")
		tmpl.Execute(w, data)
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {

}
