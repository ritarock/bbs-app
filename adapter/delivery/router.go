package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(postHandler *postHandler, commentHandler *commentHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))

	g := e.Group("/backend/api/v1")

	g.POST("/posts", postHandler.Create)
	g.GET("/posts", postHandler.GetAll)
	g.GET("/posts/:id", postHandler.GetByID)

	g.POST("/post/:post_id/comments", commentHandler.Create)
	g.GET("/post/:post_id/comments", commentHandler.GetAll)

	return e
}
