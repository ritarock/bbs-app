package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/backend/topics", readTopics)
	e.POST("/backend/topics", createTopics)
	e.GET("/backend/topics/:id", readTopic)
	e.PUT("/backend/topics/:id", updateTopic)
	e.DELETE("/backend/topics/:id", deleteTopic)

	e.GET("/backend/topics/:topic_id/comments", readComments)
	e.POST("/backend/topics/:topic_id/comments", creatComments)
	e.GET("/backend/comments/:id", readComment)
	e.PUT("/backend/comments/:id", updateComment)
	e.DELETE("/backend/comments/:id", deleteComment)
	e.Start(":8080")
}
