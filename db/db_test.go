package db

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setTestDB() string {
	return "file:ent?mode=memory&cache=shared&_fk=1"
}

func TestConnection(t *testing.T) {
	DataSource = setTestDB()
	client, gotErr := Connection()
	client.Close()
	wantErr := false
	if (gotErr != nil) != wantErr {
		t.Errorf("gotErr %v, wantErr %v", gotErr, wantErr)
	}
}
