package model

import (
	"fmt"
)

type Topic struct {
	Id     string `db:"id"`
	Title  string `db:"title"`
	Detail string `db:"detail"`
}

func ReadTopics() []Topic {
	db := connectDb()
	defer db.Close()

	topics := []Topic{}
	err := db.Select(&topics, `SELECT * FROM topics;`)
	if err != nil {
		fmt.Println(err)
	}

	return topics
}

func (t *Topic) Create() string {
	db := connectDb()
	defer db.Close()
	id := createUUId()

	_, err := db.NamedExec(
		`INSERT INTO topics (id, title, detail) VALUES (:id, :title, :detail);`,
		map[string]interface{}{
			"id":     id,
			"title":  t.Title,
			"detail": t.Detail,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	return id
}

func (t *Topic) Read() Topic {
	db := connectDb()
	defer db.Close()

	topic := Topic{}
	db.Get(&topic, `SELECT * FROM topics WHERE id = ?`, t.Id)

	return topic
}

func (t *Topic) Update() Topic {
	db := connectDb()
	defer db.Close()

	topic := Topic{
		Id:     t.Id,
		Title:  t.Title,
		Detail: t.Detail,
	}
	_, err := db.NamedExec(
		`UPDATE topics SET title = :title, detail = :detail WHERE id = :id`,
		map[string]interface{}{
			"id":     t.Id,
			"title":  t.Title,
			"detail": t.Detail,
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	return topic
}

func (t *Topic) Delete() {
	db := connectDb()
	defer db.Close()
	db.Query(`DELETE FROM topics WHERE id = ?`, t.Id)
}
