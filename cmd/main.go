package main

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ritarock/bbs-app/injector"
)

const timeout = 2 * time.Second

func main() {
	e, err := injector.InitalizeApp(timeout)
	if err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
