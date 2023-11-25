package delivery

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupRouter(
	e *echo.Echo,
	postHandler *postHandler,
	commentHandler *commentHandler,
	userHandler *userHandler,
) {
	e.POST("/backend/signup", userHandler.SignUp)
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
	auth.GET("/posts/:id", postHandler.GetById)
	auth.POST("/post/:id/comments", commentHandler.Create)
	auth.GET("/post/:id/comments", commentHandler.GetAllByPostId)
}
