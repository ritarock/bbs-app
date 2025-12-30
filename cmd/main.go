package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ritarock/bbs-app/application/usecase/auth"
	"github.com/ritarock/bbs-app/application/usecase/comment"
	"github.com/ritarock/bbs-app/application/usecase/post"
	"github.com/ritarock/bbs-app/infra/database"
	"github.com/ritarock/bbs-app/infra/handler"
	"github.com/ritarock/bbs-app/infra/handler/api"
	"github.com/ritarock/bbs-app/infra/service"
)

const dataSorce = "file:bbs.db?cache=shared&_fk=1"

func main() {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key-change-in-production"
	}
	jwtExpiration := 24 * time.Hour

	db, err := database.NewConnection(dataSorce)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	postRepo := database.NewPostRepository(db)
	commentRepo := database.NewCommentRepository(db)
	userRepo := database.NewUserRepository(db)

	passwordService := service.NewBcryptPasswordService()
	tokenService := service.NewJWTTokenService(jwtSecret, jwtExpiration)

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

	signUpUsecase := auth.NewSignUpUsecase(userRepo, passwordService, tokenService)
	signInUsecase := auth.NewSignInUsecase(userRepo, passwordService, tokenService)
	getCurrentUserUsecase := auth.NewGetCurrentUserUsecase(userRepo)

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

	authHandler := handler.NewAuthHandler(
		signUpUsecase,
		signInUsecase,
		getCurrentUserUsecase,
		tokenService,
	)

	h := handler.NewHandler(postHandler, commentHandler, authHandler)

	srv, err := api.NewServer(h, api.WithMiddleware(loging()), api.WithMiddleware(requestContext()))
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

func requestContext() middleware.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		ctx := context.WithValue(req.Context, handler.RequestContextKey(), req.Raw)
		req.SetContext(ctx)
		return next(req)
	}
}
