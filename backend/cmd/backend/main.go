package main

import (
	"bbs-app/backend/internal/db"
	"bbs-app/backend/internal/server"
)

func init() {
	db.InitDB()
}

func main() {
	server.Run()
}
