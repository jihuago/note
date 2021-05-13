package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DB()  {
	dsn := "root:123456@tcp(127.0.0.1)/fecmall"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	defer db.Close()
}
