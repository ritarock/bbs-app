package server

import (
	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()
	e.GET("/backend/topics", readTopics)
	e.POST("/backend/topics", createTopics)
	e.GET("/backend/topics/:id", readTopic)
	e.PUT("/backend/topics/:id", updateTopic)
	e.DELETE("/backend/topics/:id", deleteTopic)
	e.Start(":8080")
}
