package main

import (
	"database/sql"
	"log"
	"net/http"
	_postRepo "ritarock/bbs-app/post/repository/sqlite"
	_postUcase "ritarock/bbs-app/post/usecase"
	"time"

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

	timeoutContext := 2 * time.Second
	postucase := _postUcase.NewPostUsecase(postRepo, timeoutContext)

	handler := http.NewServeMux()
	_postDelivery.NewPostHandler(handler, postucase)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}
	server.ListenAndServe()
}
