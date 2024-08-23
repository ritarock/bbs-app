package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(postHandler *postHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	p := e.Group("/backend/api/v1")
	p.POST("/posts", postHandler.Create)
	p.GET("/posts", postHandler.GetAll)
	p.GET("/posts/:id", postHandler.GetByID)

	return e
}
