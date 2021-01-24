package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GormConnect() *gorm.DB {
	DBMS := "mysql"
	PROTOCOL := "tcp(db:3306)"
	USER := "user"
	PASS := "password"
	DBNAME := "app"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func DbInit() {
	db := GormConnect()
	defer db.Close()
	db.AutoMigrate(&Theme{})
}
