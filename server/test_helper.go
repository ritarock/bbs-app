package server

import (
	"context"
	"time"

	"github.com/ritarock/bbs-app/db"
	"github.com/ritarock/bbs-app/ent"
)

func setDBClient() *ent.Client {
	timeNow = func() time.Time {
		return time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	}()
	db.DataSource = "file:ent?mode=memory&cache=shared&_fk=1"
	client, _ := db.Connection()
	client.Schema.Create(context.Background())
	return client
}

func setupHandler() (*ent.Client, *postHandler, *commendHandler) {
	dbClient := setDBClient()

	postHandler := newPostHandler(dbClient)
	commentHandler := newCommendHandler(dbClient)

	return dbClient, postHandler, commentHandler
}

func setPost(dbClient *ent.Client) {
	dbClient.Post.
		Create().
		SetTitle("test_title").
		SetContent("test_content").
		SetPostedAt(timeNow).
		Save(context.Background())
}

func setComment(dbClient *ent.Client) {
	dbClient.Comment.
		Create().
		SetContent("test_content").
		SetPostID(1).
		SetCommentedAt(timeNow).
		Save(context.Background())
}
