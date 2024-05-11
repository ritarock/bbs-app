package delivery

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	postHandler *postHandler,
	commentHandler *commentHandler,
	userHandler *userHandler,
) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/backend/signup", userHandler.Signup)
	e.POST("/backend/login", userHandler.Login)

	auth := e.Group("/backend/api/v1")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	auth.Use(echojwt.WithConfig(config))
	auth.Use(userHandler.Session)

	auth.POST("/posts", postHandler.Create)
	auth.GET("/posts", postHandler.GetAll)
	auth.GET("/posts/:id", postHandler.GetByID)

	auth.POST("/post/:post_id/comments", commentHandler.Create)
	auth.GET("/comments", commentHandler.GetAll)
	auth.GET("/post/:post_id/comments", commentHandler.GetByPostID)
	return e
}
