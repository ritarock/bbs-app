package model

import "fmt"

type Comment struct {
	Id      string `db:"id"`
	TopicId string `db:"topic_id"`
	Body    string `db:"body"`
}

func ReadComments(topicId string) []Comment {
	db := connectDb()
	defer db.Close()

	comments := []Comment{}
	err := db.Select(&comments, `SELECT * FROM comments WHERE topic_id = ?;`, topicId)
	if err != nil {
		fmt.Println(err)
	}

	return comments
}

func (c *Comment) Create(topicId string) string {
	db := connectDb()
	defer db.Close()
	id := createUUId()

	_, err := db.NamedExec(
		`INSERT INTO comments (id, topic_id, body) VALUES (:id, :topic_id, :body);`,
		map[string]interface{}{
			"id":       id,
			"topic_id": topicId,
			"body":     c.Body,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	return id
}

func (c *Comment) Read() Comment {
	db := connectDb()
	defer db.Close()

	comment := Comment{}
	db.Get(&comment, `SELECT * FROM comments WHERE id = ?`, c.Id)

	return comment
}

func (c *Comment) Update() Comment {
	db := connectDb()
	defer db.Close()

	comment := Comment{
		Id:      c.Id,
		TopicId: c.TopicId,
		Body:    c.Body,
	}
	_, err := db.NamedExec(
		`UPDATE comments SET body = :body WHERE id = :id`,
		map[string]interface{}{
			"id":   c.Id,
			"body": c.Body,
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	return comment
}

func (c *Comment) Delete() {
	db := connectDb()
	defer db.Close()
	db.Query(`DELETE FROM comments WHERE id = ?`, c.Id)
}
