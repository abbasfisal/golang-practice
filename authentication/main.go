package main

import (
	"golang-practice/authentication/middlewares/authmiddleware"
	"golang-practice/authentication/src/controllers/admin/categorycontroller"
	"golang-practice/authentication/src/controllers/admin/productcontroller"
	"golang-practice/authentication/src/controllers/homecontroller"
	"golang-practice/authentication/src/controllers/logincontroller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homecontroller.Index)

	//admin
	http.HandleFunc("/admin/category", authmiddleware.AuthMiddleware(categorycontroller.Index))
	http.HandleFunc("/admin/product", authmiddleware.AuthMiddleware(productcontroller.Index))
	//--login
	http.HandleFunc("/admin/login", logincontroller.Index)
	http.HandleFunc("/admin/login/process", logincontroller.Process)
	//--logout
	http.HandleFunc("/admin/login/signout", logincontroller.Logout)

	log.Println("start serve on : 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("\nserver: start failed :", err.Error())
		return
	}
}
