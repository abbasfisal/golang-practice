package main

import (
	"golang-practice/crud_mysql/src/controllers/product_controller"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", product_controller.Index)
	http.HandleFunc("/product/create", product_controller.Show)
	http.HandleFunc("/product/store", product_controller.Store)
	http.HandleFunc("/product/delete", product_controller.Delete)
	http.HandleFunc("/product/edit", product_controller.Edit)
	//http.HandleFunc("/product/update", product_controller.)

	log.Println("\nstart server on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("\nserver: start fail \n ", err.Error())
	}
}
