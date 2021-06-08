package db

import (
	"data_element/app/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func DemoGorm()  {
	dsn := "homestead:secret@tcp(192.168.10.10:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	test := &models.Test{}

	db.Find(test, "id = ?", 1)
	fmt.Println(test)
}