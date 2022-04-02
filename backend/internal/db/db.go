package db

import (
	"bbs-app/backend/ent"
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DRIVER      = "sqlite3"
	DATA_SOURCE = "file:/app/data.sqlite?cache=shared&_fk=1"
)

func connection() (*ent.Client, error) {
	client, err := ent.Open(DRIVER, DATA_SOURCE)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func InitDB() {
LOOP:
	for {
		client, err := connection()
		if err != nil {
			fmt.Println("connecting db...")
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Println("connected db...")
		client.Close()
		break LOOP
	}
	client, err := connection()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
