package main

import (
	"github.com/ritarock/bbs-app/db"
	"github.com/ritarock/bbs-app/server"
)

func init() {
	db.InitDB()
}

func main() {
	// db.InitDB()
	server.Run()
}
