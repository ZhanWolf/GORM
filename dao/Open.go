package dao

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func OpenDb() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/login?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	Db = db

}
