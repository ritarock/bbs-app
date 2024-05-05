// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/labstack/echo/v4"
	"github.com/ritarock/bbs-app/adapter/delivery"
	"github.com/ritarock/bbs-app/adapter/repository"
	"github.com/ritarock/bbs-app/usecase"
	"time"
)

// Injectors from injector.go:

func InitalizeApp(timeout time.Duration) (*echo.Echo, error) {
	db := repository.NewDb()
	postRepository := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepository, timeout)
	postHandler := delivery.NewPostHandler(postUsecase)
	commentRepository := repository.NewCommentRepository(db)
	commentUsecase := usecase.NewCommentUsecase(commentRepository, timeout)
	commentHandler := delivery.NewCommentHandler(commentUsecase)
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, timeout)
	userHandler := delivery.NewUserHandler(userUsecase)
	echoEcho := delivery.NewRouter(postHandler, commentHandler, userHandler)
	return echoEcho, nil
}
