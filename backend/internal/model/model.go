package model

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	DBMS     = "mysql"
	PROTOCOL = "tcp(db:3306)"
	USER     = "user"
	PASS     = "pass"
	DBNAME   = "app"
	CONNECT  = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
)

func InitDb() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		time.Sleep(time.Second)
		return InitDb()
	}
	return db
}

func connectDb() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func createUUId() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
