package db

import (
	"context"

	"github.com/avast/retry-go"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ritarock/bbs-app/ent"
)

var DataSource = "file:data.sqlite?cache=shared&_fk=1"

func Connection() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", DataSource)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func InitDB() {
	var err error
	var client *ent.Client
	err = retry.Do(
		func() error {
			client, err = Connection()
			if err != nil {
				return err
			}
			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	err = retry.Do(
		func() error {
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
