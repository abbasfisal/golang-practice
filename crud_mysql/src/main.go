package main

import (
	"golang-practice/crud_mysql/src/controllers/product_controller"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", product_controller.Index)

	log.Println("\nstart server on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("\nserver: start fail \n ", err.Error())
	}
}
