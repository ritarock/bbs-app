package main

import (
	"database/sql"
	"log"
	_commentRepo "ritarock/bbs-app/comment/repository/sqlite"
	_commentUcase "ritarock/bbs-app/comment/usecase"
	_postRepo "ritarock/bbs-app/post/repository/sqlite"
	_postUcase "ritarock/bbs-app/post/usecase"
	"time"

	_commentDelivery "ritarock/bbs-app/comment/delivery/echo"
	_postDelivery "ritarock/bbs-app/post/delivery/echo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	commenttUcase := _commentUcase.NewCommentUsecase(commentRepo, timeoutContext)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	_postDelivery.NewPostHandler(e, postUcase)
	_commentDelivery.NewPostHandler(e, commenttUcase)

	e.Logger.Fatal(e.Start(":8080"))
}
