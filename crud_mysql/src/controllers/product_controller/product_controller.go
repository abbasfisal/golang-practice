package product_controller

import (
	"golang-practice/crud_mysql/src/entities"
	"golang-practice/crud_mysql/src/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

func Show(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("src/views/product/add.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	tmp.Execute(w, nil)
}

func Store(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var product entities.Product
	product.Name = r.Form.Get("name")
	product.Price, _ = strconv.ParseFloat(r.Form.Get("price"), 64)
	product.Quantity, _ = strconv.ParseInt(r.Form.Get("quantity"), 10, 64)
	product.Description = r.Form.Get("description")

	var pModel models.ProductModel
	pModel.Store(&product)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var pModel models.ProductModel

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return
	}
	pModel.Delete(parseID)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
