package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/backend/api/topics", getAllTopics)
	e.POST("/backend/api/topics", createTopic)
	e.GET("/backend/api/topics/:id", getTopic)
	e.PUT("/backend/api/topics/:id", updateTopic)
	e.DELETE("/backend/api/topics/:id", deleteTopic)

	e.GET("/backend/api/comments", getAllComments)
	e.POST("/backend/api/comments", createComment)
	e.GET("/backend/api/comments/:id", getComment)
	e.PUT("/backend/api/comments/:id", updateComment)
	e.DELETE("/backend/api/comments/:id", deleteComment)

	e.GET("/backend/api/topics/:topicId/comments", getCommentsByTopic)

	e.Logger.Fatal(e.Start(":8080"))
}
