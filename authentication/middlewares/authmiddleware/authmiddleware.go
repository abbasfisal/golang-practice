package authmiddleware

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func AuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "mysession")
		username := session.Values["username"]
		if username == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		f(w, r)
	}
}
