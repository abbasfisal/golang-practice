package main

import (
	"golang-practice/authentication/src/controllers/admin/categorycontroller"
	"golang-practice/authentication/src/controllers/admin/productcontroller"
	"golang-practice/authentication/src/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homecontroller.Index)

	//admin
	http.HandleFunc("/admin/category", categorycontroller.Index)
	http.HandleFunc("/admin/product", productcontroller.Index)

	log.Println("start serve on : 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("\nserver: start failed :", err.Error())
		return
	}
}
