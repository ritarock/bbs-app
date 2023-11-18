package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ritarock/bbs-app/adapter/delivery"
	"github.com/ritarock/bbs-app/adapter/repository"
	"github.com/ritarock/bbs-app/usecase"
)

const dataSource = "file:data.sqlite?cache=shared&_fk=1"
const timeout = 2 * time.Second

func main() {
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	postRepo := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepo, timeout)

	commentRepo := repository.NewCommentRepository(db)
	commentUsecase := usecase.NewCommentUsecase(commentRepo, timeout)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
	}))

	delivery.NewPostHandler(e, postUsecase)
	delivery.NewCommentHandler(e, commentUsecase)

	e.Logger.Fatal(e.Start(":8080"))
}
