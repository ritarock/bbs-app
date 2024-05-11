package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ritarock/bbs-app/injector"
	"go.uber.org/zap"
)

const timeout = 2 * time.Second

func main() {
	e, err := injector.InitalizeApp(timeout)
	if err != nil {
		log.Fatal(err)
	}

	logger, _ := zap.NewProduction()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	e.Logger.Fatal(e.Start(":8080"))
}
