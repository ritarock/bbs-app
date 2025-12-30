package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ritarock/bbs-app/application/usecase/comment"
	"github.com/ritarock/bbs-app/application/usecase/post"
	"github.com/ritarock/bbs-app/infra/database"
	"github.com/ritarock/bbs-app/infra/handler"
	"github.com/ritarock/bbs-app/infra/handler/api"
)

const dataSorce = "file:bbs.db?cache=shared&_fk=1"

func main() {
	db, err := database.NewConnection(dataSorce)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	postRepo := database.NewPostRepository(db)
	commentRepo := database.NewCommentRepository(db)

	createPostUsecase := post.NewCreatePostUsecase(postRepo)
	getPostUsecase := post.NewGetPostUsecase(postRepo)
	listPostUsecase := post.NewListPostUsecase(postRepo)
	updatePostUsecase := post.NewUpdatePostUsecase(postRepo)
	deletePostUsecase := post.NewDeletePostUsecase(postRepo)

	createCommentUsecase := comment.NewCreateCommentUsecase(commentRepo)
	getCommentUsecase := comment.NewGetCommentUsecase(commentRepo)
	listCommentsUsecase := comment.NewListCommentsUsecase(commentRepo)
	updateCommentUsecase := comment.NewUpdateCommentUsecase(commentRepo)
	deleteCommentUsecase := comment.NewDeleteCommentUsecase(commentRepo)

	postHandler := handler.NewPostHandler(
		createPostUsecase,
		getPostUsecase,
		listPostUsecase,
		updatePostUsecase,
		deletePostUsecase,
	)

	commentHandler := handler.NewCommentHandler(
		createCommentUsecase,
		getCommentUsecase,
		listCommentsUsecase,
		updateCommentUsecase,
		deleteCommentUsecase,
	)

	h := handler.NewHandler(postHandler, commentHandler)

	srv, err := api.NewServer(h, api.WithMiddleware(loging()))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}

func loging() middleware.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		start := time.Now()
		resp, err := next(req)
		duration := time.Since(start)

		log.Printf("%s %s %s", req.Raw.Method, req.Raw.URL.Path, duration)

		return resp, err
	}
}
