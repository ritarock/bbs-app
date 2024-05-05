//go:build wireinject
// +build wireinject

package injector

import (
	"time"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/ritarock/bbs-app/adapter/delivery"
	"github.com/ritarock/bbs-app/adapter/repository"
	"github.com/ritarock/bbs-app/usecase"
)

func InitalizeApp(timeout time.Duration) (*echo.Echo, error) {
	wire.Build(
		repository.NewDb,
		repository.NewPostRepository,
		repository.NewCommentRepository,
		repository.NewUserRepository,

		usecase.NewPostUsecase,
		usecase.NewCommentUsecase,
		usecase.NewUserUsecase,

		delivery.NewPostHandler,
		delivery.NewCommentHandler,
		delivery.NewUserHandler,
		delivery.NewRouter,
	)
	return nil, nil
}
