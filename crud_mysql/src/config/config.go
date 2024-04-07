package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func GetDb() (db *sql.DB, err error) {

	db, err = sql.Open("mysql", "root:root@/gopractice")
	if err != nil {
		log.Fatal("\nDb : connection failed \n ", err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return
}
