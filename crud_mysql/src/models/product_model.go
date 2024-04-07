package models

import (
	"golang-practice/crud_mysql/src/config"
	"golang-practice/crud_mysql/src/entities"
)

type ProductModel struct {
}

func (*ProductModel) FindAll() ([]entities.Product, error) {
	db, err := config.GetDb()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from products ")
	if err != nil {
		return nil, err
	}

	var products []entities.Product
	for rows.Next() {
		var p entities.Product
		rows.Scan(&p.Id, &p.Name, &p.Price, &p.Quantity, &p.Description)

		products = append(products, p)
	}
	return products, nil
}
