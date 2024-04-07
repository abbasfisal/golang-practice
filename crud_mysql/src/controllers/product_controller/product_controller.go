package product_controller

import (
	"golang-practice/crud_mysql/src/models"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var pm models.ProductModel

	products, _ := pm.FindAll()

	data := map[string]interface{}{
		"products": products,
	}
	tmp, err := template.ParseFiles("src/views/product/index.html")
	if err != nil {
		log.Fatal("\nfail: parse view product/index \n", err.Error())
	}
	tmp.Execute(w, data)
}
