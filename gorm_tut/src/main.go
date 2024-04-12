package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gopractice?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("[error] Db Opening : ", err)
	}

	log.Println("\n[info] Db: connected successfully ")

	db.AutoMigrate(Teacher{}, &Blog{}, &Comment{})
	//var t Teacher
	var tt []Teacher
	res := db.Debug().Find(&tt, []int{13, 14, 15})
	fmt.Println(tt)

	errors.Is(res.Error, gorm.ErrRecordNotFound)
}

type Teacher struct {
	gorm.Model
	UserName  string
	FirstName string `gorm:"default:yaali"`
}

func (t *Teacher) BeforeCreate(tx *gorm.DB) error {
	t.UserName = t.UserName + " -teacher"
	return nil
}

type Blog struct {
	gorm.Model
	Title       string
	Description *string
}
type Comment struct {
	gorm.Model
	Text string
}
