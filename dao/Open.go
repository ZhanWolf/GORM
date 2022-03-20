package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func OpenDb() error {
	db, err := sql.Open("mysql", "root:123456@/login")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return err
	}
	Db = db

	return nil

}
