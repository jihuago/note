package controllers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	id int
	name string
}

func initDB() (err error){
	dsn := "user:password@tcp(192.168.10.10:3306)/yii_test?charset=utf8mb4&parseTime=True"

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10)

	return
}

func QueryRowDemo()  {
	err := initDB()
	if err != nil {
		panic(err)
	}
	sqlStr := "select id, name from go_user where id = ?"

	var u user
	err = db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err%v\n", err)
		return
	}

	fmt.Printf("id:%d name:%s age:%d\n",)

}
