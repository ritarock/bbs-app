package db

import (
	"context"

	"github.com/avast/retry-go"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ritarock/bbs-app/ent"
)

const (
	DRIVER      = "sqlite3"
	DATA_SOURCE = "file:data.sqlite?cache=shared&_fk=1"
)

func Connection() (*ent.Client, error) {
	client, err := ent.Open(DRIVER, DATA_SOURCE)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func InitDB() {
	err := retry.Do(
		func() error {
			client, err := Connection()
			if err != nil {
				return err
			}
			defer client.Close()

			err = client.Schema.Create(context.Background())
			if err != nil {
				return err
			}
			return nil
		},
	)
	if err != nil {
		panic(err)
	}
}
