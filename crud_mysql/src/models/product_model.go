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
func (*ProductModel) Store(product *entities.Product) bool {
	db, err := config.GetDb()
	if err != nil {
		return false
	}
	result, ExeErr := db.Exec("insert into products (name , price , quantity , description) values (? , ? , ? , ? )",
		product.Name,
		product.Price,
		product.Quantity,
		product.Description,
	)
	if ExeErr != nil {
		return false
	}
	rowsAffected, Aerr := result.RowsAffected()
	if Aerr != nil {
		return false
	}
	return rowsAffected > 0
}

func (*ProductModel) Delete(id int64) bool {
	db, err := config.GetDb()
	if err != nil {
		return false
	}
	res, Eerr := db.Exec("delete from products where id = ? ", id)
	if Eerr != nil {
		return false
	}
	rowsAffected, AfErr := res.RowsAffected()
	if AfErr != nil {
		return false
	}
	return rowsAffected > 0
}
