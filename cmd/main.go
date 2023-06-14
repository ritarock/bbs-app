package main

import (
	"database/sql"
	"log"
	"net/http"
	_commentRepo "ritarock/bbs-app/comment/repository/sqlite"
	_commentUcase "ritarock/bbs-app/comment/usecase"
	_postRepo "ritarock/bbs-app/post/repository/sqlite"
	_postUcase "ritarock/bbs-app/post/usecase"
	"time"

	_commentDelivery "ritarock/bbs-app/comment/delivery/http"
	_postDelivery "ritarock/bbs-app/post/delivery/http"

	_ "github.com/mattn/go-sqlite3"
)

var dataSource = "file:data.sqlite?cache=shared&_fk=1"

func main() {
	conn, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	postRepo := _postRepo.NewSqlitePostRepository(conn)
	commentRepo := _commentRepo.NewsqliteCommentRepository(conn)

	timeoutContext := 2 * time.Second
	postUcase := _postUcase.NewPostUsecase(postRepo, timeoutContext)
	commentUcase := _commentUcase.NewCommentUsecase(commentRepo, timeoutContext)

	handler := http.NewServeMux()
	_postDelivery.NewPostHandler(handler, postUcase)
	_commentDelivery.NewCommentHandler(handler, commentUcase)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}
	server.ListenAndServe()
}
